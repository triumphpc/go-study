// io.Pipe в Go используется для создания синхронного канала в памяти, который соединяет io.Reader и io.Writer.
//Это позволяет передавать данные между двумя потоками или горутинами без использования промежуточного буфера. io.Pipe
//полезен, когда вы хотите напрямую передавать данные из одной части программы в другую, например, из функции, которая
//генерирует данные, в функцию, которая их обрабатывает.
//
//Как работает io.Pipe
//io.PipeReader: Читает данные, которые были записаны в io.PipeWriter.
//io.PipeWriter: Записывает данные, которые могут быть прочитаны io.PipeReader.
//Поскольку io.Pipe является синхронным, операции чтения и записи блокируются до тех пор, пока другая сторона не будет готова. Это делает его подходящим для потоковой передачи данных между горутинами.

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Ride - структура, представляющая информацию о поездке.
type Ride struct {
	Time     time.Time
	Distance float64
	Price    float64
}

// Conn - структура, представляющая соединение с базой данных.
type Conn struct{}

// QueryRidesIn - возвращает канал с поездками для указанного местоположения.
// Закрывает канал, когда данные заканчиваются.
func (c *Conn) QueryRidesIn(location string) <-chan Ride {
	ch := make(chan Ride)

	// Симулируем запрос к базе данных и отправку данных в канал
	go func() {
		defer close(ch)
		// Симулируем несколько поездок
		for i := 0; i < 5; i++ {
			ch <- Ride{
				Time:     time.Now(),
				Distance: float64(i) * 10,
				Price:    float64(i) * 5,
			}
			time.Sleep(500 * time.Millisecond) // Симулируем задержку
		}
	}()

	return ch
}

// encodeRides - кодирует поездки из канала в writer
func encodeRides(ch <-chan Ride, w io.WriteCloser) error {
	enc := json.NewEncoder(w)
	defer w.Close() // Закрываем writer, чтобы сигнализировать об окончании данных
	for r := range ch {
		if err := enc.Encode(r); err != nil {
			return err
		}
	}
	return nil
}

// queryHandler - обработчик HTTP-запросов для запросов поездок
func queryHandler(w http.ResponseWriter, r *http.Request) {
	const maxSize = 1024 // Ограничение на размер тела запроса

	data, err := io.ReadAll(io.LimitReader(r.Body, maxSize))
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	location := string(data)

	conn := &Conn{} // Симулируем соединение с базой данных
	ch := conn.QueryRidesIn(location)

	// Создаем pipe для соединения кодировщика с HTTP-ответом
	rp, wp, err := os.Pipe()
	if err != nil {
		http.Error(w, "can't create pipe", http.StatusInternalServerError)
		return
	}

	// Запускаем горутину для кодирования поездок в writer pipe
	go func() {
		if err := encodeRides(ch, wp); err != nil {
			log.Printf("error: can't encode: %s", err)
		}
	}()

	// Копируем данные из reader pipe в HTTP-ответ
	_, err = io.Copy(w, rp)
	if err != nil {
		log.Printf("error: can't copy to response: %s", err)
	}
}

func main() {
	http.HandleFunc("/query", queryHandler)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
