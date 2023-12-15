package main

import (
	"fmt"
	"time"
)

type ModelDataResponseBody struct {
	Path      string `json:"path,omitempty"`
	Status    string `json:"status,omitempty"`
	NeedCrypt bool   `json:"needCrypt,omitempty"`
	Test      map[string]string
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

	testXXX(nil)

}

func testXXX(mm map[string]interface{}) {
	if mm == nil {
		fmt.Println("V")
		return
	}

	fmt.Println("XX")

}
