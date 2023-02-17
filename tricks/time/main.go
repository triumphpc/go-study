package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"time"
)

type Type1 map[string]interface{}
type Type2 map[string]interface{}

type View struct {
	ID        string          `json:"id"`
	GUID      string          `json:"guid,omitempty"`
	FileID    string          `json:"file_id,omitempty"`
	ServiceID string          `json:"service,omitempty"`
	Code      string          `json:"pin,omitempty"`
	Created   time.Time       `json:"created"`
	Updated   *time.Time      `json:"updated,omitempty"`
	Reserve   *time.Time      `json:"reserve_time,omitempty"`
	Extra     json.RawMessage `json:"extra,omitempty"`
}

func main() {

	str := "ООО \"АСТРАЛ-СОФТ\" \"SOFT Coom\" и КО_МП  "

	sanitizeField(&str)

	fmt.Println(str)

	//strF := strings.ReplaceAll(str, "\"", "\\\"")
	//fmt.Println(strF)

	//if str == "" || utf8.RuneCountInString(str) > 255 {
	//	fmt.Println("VVV")
	//}

	//regex := regexp.MustCompile("[a-z]+")
	//out := regex.ReplaceAllString(str, "")
	//
	//fmt.Println(out)
	//
	//checkSymbols := `[а-я]+`
	//matched, _ := regexp.MatchString(checkSymbols, out)
	//fmt.Println(matched) // true

	//tt := time.Duration(time.Second * 3)

	//mm := make([]string, 2, 3)

	//n := 5
	//time.Sleep(time.Duration(n) * time.Second)
	//
	//fmt.Println("VVV")

	//data := Type1{
	//	"test1": "RRRRR",
	//	"test2": 12345,
	//	"test3": "xxxx",
	//}
	//
	//fmt.Printf("%T", data)
	//
	//data2 := Type2(data)
	//fmt.Printf("%T", data2)
	//fmt.Printf("%v", data2)

	//data2 :=

	//t := time.Now()
	//
	//time.Sleep(2 * time.Second)
	//
	//t2 := time.Now()
	//
	//fmt.Println(t.After(t2))
	//fmt.Println(t.Sub(t2))

	//t3 := time.NewTimer(0)
	//
	//for {
	//	t3.Reset(time.Second * 3)
	//
	//	select {
	//	case <-t3.C:
	//		fmt.Println("Yep")
	//	}
	//}

	//format := "02.01.2006 15:04:05 Z07:00"
	//tm, _ := time.Parse(format, "17.08.2023 10:51:58 +06:00")
	//
	//fmt.Println(tm)
	//nano := tm.UnixNano()
	//fmt.Println(nano) //1692258717000000000
	//
	//tm2 := time.Unix(0, nano)
	//fmt.Println(tm2)
	//
	//tm3 := DateTimeRFC3339{Time: tm2.UTC()}.String()
	//fmt.Println(tm3)

}

func testErr() error {

	var err error
	defer func() {

		if err != nil {
			fmt.Println("Зашли сюда")
		}

		return

	}()

	err = errors.New("TEST")

	return err

}

// sanitizeField - подготовка поля для дальнейшей обработки
// С фронта могут прийти "битые" поля, тут их приводим в нормальный вид
func sanitizeField(s *string) {
	// Проверим, что вообще поле заполнено и есть символы кириллицы
	// В противном случае отдаем пустую строку
	rxp := regexp.MustCompile("[^а-яА-Я]+")
	tmp := rxp.ReplaceAllString(*s, "")
	if tmp == "" {
		*s = ""

		return
	}
	// Любая кавычка заменяется на стандартную двойную
	rxp = regexp.MustCompile("['<>`„‹’”›«»]")
	*s = rxp.ReplaceAllString(*s, "\"")

	// Удаляем не русские символы
	rxp = regexp.MustCompile("[a-zA-Z,]")
	*s = rxp.ReplaceAllString(*s, "")

	//rxp = regexp.MustCompile("[A-Z]")
	//*s = rxp.ReplaceAllString(*s, "")

	// Удаляем двойные пробелы
	rxp = regexp.MustCompile("\\s+")
	*s = rxp.ReplaceAllString(*s, " ")

	// Удаляем все повторные кавычки
	rxp = regexp.MustCompile("\"+")
	*s = rxp.ReplaceAllString(*s, "\"")

	// Утраиваем двойные кавычки, чтобы корректно сгенерировались в cryptcp
	// Формат  CN=\"\"\"АСТРАЛ\"\"\" кажет "ООО "АСТРАЛ" в серте после генерации
	rxp = regexp.MustCompile("\"")
	*s = rxp.ReplaceAllString(*s, "\"\"\"")
}

// NewDateTimeRFC3339 - конструктор даты в формате RFC3339
func NewDateTimeRFC3339(t time.Time) DateTimeRFC3339 {
	return DateTimeRFC3339{
		Time: t,
	}
}

// DateTimeRFC3339 - дата и время в формате RFC3339 "2006-01-02T15:04:05Z07:00"
type DateTimeRFC3339 struct {
	time.Time
}

func (d *DateTimeRFC3339) UnmarshalJSON(b []byte) (err error) {
	var s string

	if err = json.Unmarshal(b, &s); err != nil {
		return err
	}

	d.Time, err = time.Parse(time.RFC3339, s)
	if err != nil {
		// Ядро отчетности иногда присылает RFC3339 без указания таймзоны, они это считают Z
		var e error
		s += "Z"

		if d.Time, e = time.Parse(time.RFC3339, s); e != nil {
			return err
		}
		return nil
	}

	return err
}

func (d DateTimeRFC3339) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format(time.RFC3339))

}

func (d DateTimeRFC3339) String() string {
	return d.Format(time.RFC3339)
}
