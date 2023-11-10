package main

import (
	"fmt"
	"strings"
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

	str := "123 12"

	nnn := strings.ReplaceAll(str, " ", "")
	ss := nnn[5:6]

	fmt.Println(ss)

	//fmt.Println(res)

}
