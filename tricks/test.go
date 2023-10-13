package main

import (
	"encoding/json"
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
	str := "{\n    \"path\": \"/app-web/v1/task-async/status/3c462d0d-c90d-438c-a680-59c22e87a9fd\",\n    \"requestId\": \"82f05888-1524-44bc-92f9-12cafd011dfa\",\n    \"timeStamp\": \"2023-10-09T11:54:35.371+0300\",\n    \"data\": {\n        \"path\": \"https://adp-oapi-redis.gnivc.ru/webdav/ens/2023/10/09/9675324089_3c462d0d-c90d-438c-a680-59c22e87a9fd.czip\",\n        \"retryCount\": 0,\n        \"needCrypt\": true,\n        \"creationRequestId\": \"dd4eaad6-bda9-4b89-8674-28d7abc3f464\"\n    }\n}"

	res := new(ModelDataResponse)
	_ = json.Unmarshal([]byte(str), res)

	fmt.Println(res)

}
