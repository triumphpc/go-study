package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"syscall"
)

// Location структура для хранения данных о местоположении
type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Name      string  `json:"name"`
}

func main() {
	// Открываем файл для чтения
	file, err := os.OpenFile("data.json", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Получаем информацию о файле
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}

	// Создаем memory mapping файла
	data, err := syscall.Mmap(
		int(file.Fd()),       // файловый дескриптор
		0,                    // смещение
		int(fileInfo.Size()), // размер
		syscall.PROT_READ,    // права доступа (только чтение)
		syscall.MAP_SHARED,   // тип отображения
	)
	if err != nil {
		log.Fatalf("Failed to mmap: %v", err)
	}
	defer syscall.Munmap(data) // освобождаем memory mapping при завершении

	// Текущая позиция в данных
	pos := 0

	// Префикс для поиска локаций
	locPrefix := []byte("loc:{")

	// Структура для хранения распарсенной локации
	var loc Location

	// Основной цикл обработки
	for {
		// Ищем следующее вхождение "loc:" в данных
		i := bytes.Index(data[pos:], locPrefix)
		if i == -1 {
			// Больше локаций не найдено
			break
		}

		// Перемещаемся за префикс
		i += len(locPrefix) - 1
		start := pos + i

		// Ищем закрывающую скобку текущего JSON-объекта
		size := bytes.IndexByte(data[start:], '}')
		if size == -1 {
			// Если не найдена закрывающая скобка, прерываем
			break
		}

		// Включаем закрывающую скобку
		size++

		// Парсим JSON в структуру Location
		if err := json.Unmarshal(data[start:start+size], &loc); err != nil {
			log.Fatalf("Error parsing JSON: %s", err)
		}

		// Выводим распарсенную локацию
		fmt.Printf("Found location: %+v\n", loc)

		// Перемещаем указатель за текущий JSON
		pos = start + size + 1
	}
}
