package main

import (
	"fmt"
	"time"
)

type ModelDataResponseBody struct {
	Path      string `json:"path,omitempty"`
	Status    string `json:"status,omitempty"`
	NeedCrypt bool   `json:"needCrypt,omitempty"`
}

type ModelError struct {
	Code      string `json:"code"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

type ModelDataResponse struct {
	Path      string                 `json:"path"`
	RequestID string                 `json:"requestId"`
	TimeStamp time.Time              `json:"timeStamp"`
	Data      *ModelDataResponseBody `json:"data,omitempty"`
	Error     *ModelError            `json:"error,omitempty"`
}

func main() {
	t := time.Now()
	year, month, day := t.Date()
	st := time.Date(year, month, day-2, 0, 0, 0, 0, t.Location())

	fmt.Println(st.Unix())

}
