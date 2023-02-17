package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type EDOFileImport struct {
	AbonentID string      `json:"ИдАбонента"`
	SignDss   *EDOSignDss `json:"ПараметрыПодписиDSS,omitempty"`
}

type EDOSignDss struct {
	CertID     string `json:"ИдСертификата"`
	WithNotify *bool  `json:"ОтправитьУведомление,omitempty"`
}

type coreMessage struct {
	TraceID string `json:"traceId"`
	Data    []byte `json:"data"`
}

type DateTimeRFC3339WithoutColon struct {
	time.Time
}

const timeZoneSeparator = "+"
const ensDateTimeLayout = "2006-01-02T15:04:05Z0700"

func (d *DateTimeRFC3339WithoutColon) UnmarshalJSON(b []byte) (err error) {
	var s string

	if err = json.Unmarshal(b, &s); err != nil {
		return err
	}

	strLen := len(s)
	zIndex := strings.Index(s, timeZoneSeparator)
	if zIndex == -1 || strings.Contains(s[zIndex:strLen], ":") {
		d.Time, err = time.Parse(time.RFC3339, s)
	} else {
		d.Time, err = time.Parse(ensDateTimeLayout, s)
	}

	return err
}

func (d DateTimeRFC3339WithoutColon) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format(time.RFC3339))
}

type modelDelayedTask struct {
	Path      string                      `json:"path"`
	RequestID string                      `json:"requestId"`
	TimeStamp DateTimeRFC3339WithoutColon `json:"timeStamp"`
	Data      string                      `json:"data"`
	Error     *modelError                 `json:"error,omitempty"`
	//Error string `json:"error,omitempty"`
}

type modelDelayedTaskServer struct {
	Path      string                      `json:"path"`
	RequestID string                      `json:"requestId"`
	TimeStamp DateTimeRFC3339WithoutColon `json:"timestamp"`
	Status    int                         `json:"status"`
	Message   string                      `json:"message"`
	Error     string                      `json:"error"`
}

type modelError struct {
	Code      string `json:"code"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

type MyType struct {
	st string
}

type SomeType interface {
}

func test(req interface{}) {

	var dd *MyType
	switch val := req.(type) {
	case *MyType:
		dd = val
	default:
		fmt.Println("vvvv")
	}

	fmt.Println(dd)

}
func main() {

	//js := []byte("{\"timestamp\":\"2023-01-19T09:34:51.019+03:00\",\"path\":\"/app-web/v1/handbooks/id/kbk/delayed\",\"status\":400,\"error\":\"openApi.badAuthenticationSchema\",\"message\":\"В заголовке 'Authorization' указана неправильная схема аутентификации. Должна быть указана схема аутентификации 'Bearer '.\",\"requestId\":\"4a928f24-58ee-46cb-aeed-2291db2cc1a8\"}")
	//
	res := new(MyType)
	res.st = "fdfdfdf"

	test(res)

	//
	//err := json.Unmarshal(js, res)
	//
	//fmt.Println(err, res)

	//ss := []string{"1", "2", "3"}
	//
	//ss = ss[1:]
	//
	//fmt.Println(ss)

	//now := time.Now()
	//yyyy, mm, dd := now.Date()
	//
	////Обновляем каждый день в 2:30 ночи
	//updateTime := time.Date(yyyy, mm, dd, 2, 30, 0, 0, now.Location())
	//
	//fmt.Println(updateTime)

	//
	//res, err := json.Marshal(&coreMessage{
	//	TraceID: "0b6804fa-d3b5-44a8-84be-a259cceb60eb",
	//	Data:    []byte("eyJwYXlsb2FkIjp7ImRlc2NyaXB0aW9uIjoiIiwiZGF0YSI6eyJkb2NmbG93SWQiOiIxZDFlOTU1Ni03ZmJhLTQ5ZDQtODJiMS1kNzZhZDFmYmU3YTciLCJhYm9uZW50SWQiOiIyZjBkNjg0OS1kYWYyLTRmMmMtYjQyZi00YjAzMjc5NzhhYjMiLCJnYXRlSWQiOiI1MTU5RTZBOS0zNjNBLTQxRjgtQTk2Ny0yNUI3MzZBQUQzNkQiLCJ4UmVhbElwIjoiOTIuMTYuNTQuMTI0IiwicmVjZWl2ZXJUeXBlIjoxLCJyZWNlaXZlciI6Ijk5OTkiLCJ0aGVtZSI6ItCi0LXRgdGC0L7QstCw0Y8g0YLQtdC80LAyMyIsIm1lc3NhZ2UiOiLQotC10YHRgtC+0LLQvtC1INGB0L7QvtCx0YnQtdC90LjQtTIzIiwib3BlcmF0aW9uRGF0YSI6IiIsImZpbGVzIjpbXSwicGFja2FnZUhlYWRlcnMiOlt7ImtleSI6ItCS0LXRgNGB0J/RgNC+0LMiLCJ2YWx1ZSI6ItCQ0YHRgtGA0LDQuy7Qn9C70LDRgtGE0L7RgNC80LAifSx7ImtleSI6ItCY0YHRgtC+0YfQvdC40LoiLCJ2YWx1ZSI6IkFQSSJ9XX0sIm1lc3NhZ2VJZCI6ImRkODAxYWRkLTk2MjAtNGIyYS04ZDY4LWQyY2ZhMjhkMmYwMCIsInR5cGUiOjgsImVycm9yIjoiIiwib3BlcmF0aW9uRGF0YSI6IiJ9LCJwYXlsb2FkVHlwZSI6IiIsInZlcnNpb24iOiIiLCJzZW5kZXIiOiIiLCJjcmVhdGVkQXQiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiJ9"),
	//})
	//
	//fmt.Println(err)
	//
	//data := new(coreMessage)
	//err = json.Unmarshal(res, data)
	//
	//fmt.Println(err)
	//fmt.Println(data)

	//res := 1001
	//
	//fmt.Println(res % 1000)
	//arg := new(EDOFileImport)
	//data := `{"ИдАбонента":"3", "ПараметрыПодписиDSS":{"ИдСертификата": "444", "ОтправитьУведомление": true}}`
	//
	//if err := json.Unmarshal([]byte(data), arg); err != nil {
	//}
	//
	//if arg.SignDss != nil {
	//	fmt.Println(arg.SignDss.CertID)
	//	if arg.SignDss.WithNotify != nil {
	//		fmt.Println(*arg.SignDss.WithNotify)
	//
	//	}
	//}
	//
	//fmt.Printf("%", arg)
	//
	//tt := false
	//
	//fmt.Printf("%t", tt)
	//
	//ff := []byte("ccccccc")
	//cccc := []byte("gggggg")
	//var content []byte
	//
	//content = append(ff, []byte("\n")...)
	//content = append(content, cccc...)
	//
	//fmt.Println(string(content))

}
