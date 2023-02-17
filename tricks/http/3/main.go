package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"
)

var rm sync.RWMutex
var cache int = 0

type modelRejectDoc struct {
	AbonentGlobalID     string `json:"abonentGlobalId"`
	DocflowID           string `json:"docflowId"`
	Message             string `json:"message"`
	IsWriteMarkingCodes *bool  `json:"isWriteMarkingCodes,omitempty"`
}

type testStr struct {
	here *bool
}

type tokenForCertIDResponse struct {
	// Поля синхронного запроса
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Content string `json:"Content,omitempty"`
	// Поля асинхронного запроса
	TransactionID string `json:"TransactionId,omitempty"`
}

type modelDataResponse struct {
	Path      string                 `json:"path"`
	RequestID string                 `json:"requestId"`
	TimeStamp ensDateTimeRFC3339     `json:"timeStamp"`
	Data      *modelDataResponseBody `json:"data,omitempty"`
	Error     *modelDelayedTaskError `json:"error,omitempty"`
}

type modelDataResponseBody struct {
	Path   string `json:"path,omitempty"`
	Status string `json:"status,omitempty"`
}

type modelDelayedTaskError struct {
	Code      string `json:"code"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

func main() {
	from, _ := time.Parse(time.RFC3339, "2022-12-08T09:48:58Z")

	fmt.Println(from.String())

	//str := "https://adp-oapi-redis.gnivc.ru/webdav/ens/2022/12/08/9675324089_2b54a53e-729c-4adf-9cce-a8f3d63b772e.czip"
	//parts, err := url.Parse(str)
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(parts.Path)
	//fmt.Println(parts.RawQuery)

	//str := "{\n    \"path\": \"/app-web/v1/task-async/status/17b25a6c-8b34-4df1-bf33-0f9f7120996d\",\n    \"requestId\": \"25f99626-5816-4819-8cde-bc180a7e0b5c\",\n    \"timeStamp\": \"2022-12-07T14:56:37.573+0000\",\n    \"data\": {\n        \"path\": \"https://adp-oapi-redis.gnivc.ru/webdav/ens/2022/12/07/9675324089_17b25a6c-8b34-4df1-bf33-0f9f7120996d.czip\"\n    }\n}"

	//t := "http://10.164.65.101:8111/2022/06/12/1234567890_f0323e54-4bce-4aa5-a815-bbc57f5d0fa4.zip"
	//v, err := url.Parse(t)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(v.Path)
	//fmt.Println(v.Scheme)
	//fmt.Println(v.ForceQuery)
	//fmt.Println(v.Fragment)
	//fmt.Println(v.RawPath)
	//fmt.Println(v.RawQuery)
	//fmt.Println(v.Host)

	//
	//str := []byte("{\"path\":\"/app-web/v1/task-async/status/66a4a43b-21b1-4a51-9dc0-32a5a3e89f21\",\"requestId\":\"552854ab-41a0-4e74-afd9-7ef372036bdb\",\"timeStamp\":\"2022-12-08T06:40:07.941+0000\",\"data\":{\"path\":\"https://adp-oapi-redis.gnivc.ru/webdav/ens/2022/12/08/9675324089_66a4a43b-21b1-4a51-9dc0-32a5a3e89f21.czip\"}}")
	//
	//stringReader := strings.NewReader(str)
	//stringReadCloser := ioutil.NopCloser(stringReader)

	//resp := new(modelDataResponse)

	//err := json.Unmarshal(str, resp)
	//if err := json.NewDecoder(stringReadCloser).Decode(resp); err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(resp)
	//fmt.Println(err)

	//
	//var w *sync.WaitGroup
	//
	//w.Wait()

	//nt := new(testStr)

	//rr := false
	//nt = &testStr{
	//	here: &rr,
	//}

	//testFuncr(nt.here)

	//tt := time.Now()
	////
	//fmt.Printf("вот тут %v", tt.UTC())
	//fmt.Printf("вот тут %v", tt.Local())

	//func() {
	//	defer test()
	//
	//	fmt.Println("here")
	//
	//}()

	//return
	//
	//dur := time.Duration(1 * time.Second)
	//
	//fmt.Printf("Отложен %v", dur)

	//now := time.Now().Add(dur)

	//fmt.Println(now.Local())

	//for i := 0; i < 5; i++ {
	//	go func(m int) {
	//		fmt.Printf("Старт %d\n", m)
	//		for {
	//			read()
	//		}
	//	}(i)
	//
	//}
	//
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	//defer func() {
	//	//	cache = 0
	//	//}()
	//
	//	fmt.Println(read())
	//
	//	write()
	//})
	//
	//_ = http.ListenAndServe(":3333", nil)
}

func doIt(rc io.ReadCloser) {
	defer rc.Close()
	buf := make([]byte, 4)
	n, err := rc.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf[:n])
}

func testFuncr(test *bool) {
	nt := new(modelRejectDoc)

	nt.IsWriteMarkingCodes = test

	data, err := json.Marshal(nt)

	fmt.Println(data, err)

	var result222 = new(modelRejectDoc)
	err = json.Unmarshal(data, result222)
	fmt.Println(err)
	//fmt.Println(*result222.IsWriteMarkingCodes)

}

func test() {
	fmt.Println("MMMM")

}

func read() int {
	rm.RLock()
	defer rm.RUnlock()

	return cache

}

func write() {
	rm.Lock()
	defer rm.Unlock()

	cache = cache + 1

}

/*
ensDateTimeRFC3339 - дата и время в формате RFC3339 "2006-01-02T15:04:05Z07:00"
или "2006-01-02T15:04:05Z0700" (во временной зоне нет двоеточия)
*/
type ensDateTimeRFC3339 struct {
	time.Time
}

const timeZoneSeparator = "+"
const ensDateTimeLayout = "2006-01-02T15:04:05Z0700"

func (d *ensDateTimeRFC3339) UnmarshalJSON(b []byte) (err error) {
	var s string

	if err = json.Unmarshal(b, &s); err != nil {
		return err
	}

	strLen := len(s)
	zIndex := strings.Index(s, timeZoneSeparator)
	log.Println(strLen)
	log.Println(zIndex)
	if zIndex == -1 || strings.Contains(s[zIndex:strLen], ":") {
		log.Println(1)
		d.Time, err = time.Parse(time.RFC3339, s)
	} else {
		log.Println(2)
		d.Time, err = time.Parse(ensDateTimeLayout, s)
	}

	return err
}

func (d ensDateTimeRFC3339) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format(time.RFC3339))
}
