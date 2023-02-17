package main

import (
	"bytes"
	"fmt"
	"time"
)

type PartnerIDs struct {
	IDs []string
}

const tt = time.Time(time.Second)

func main() {

	//hd := make(map[string]string, 3)
	//
	//hd["test"] = "fdfdf"
	//
	////_, ok := hd["tegst"]
	//val, ok := hd["testf"]
	//
	//fmt.Println(val, ok)

	//ids := make([]string, 0, 3)
	//ids = append(ids, "one", "two", "three")
	//
	//var buf []byte
	//var err error
	//if buf, err = json.Marshal(ids); err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(buf)

	//var partners *PartnerIDs

	//list := make([]string, 0, 10)
	//![](../../../../../../../../../../../var/folders/5x/c5lk98_j31v5p3jh03g8pq2h0000gn/T/TemporaryItems/NSIRD_screencaptureui_7VAjja/Screenshot 2023-02-17 at 15.23.27.png)
	//r := bytes.NewReader(buf)
	//
	////fmt.Println(r.Read())
	//
	//if err = json.NewDecoder(r).Decode(&list); err != nil {
	//	fmt.Println(err)
	//
	//	return
	//}
	//
	//fmt.Println(list)

	//data := make([]string, 0, 10)
	//
	//data = append(data, "fdfd")
	//
	//tt := func(ads ...string) {
	//
	//	for idx := range ads {
	//		fmt.Println(ads[idx])
	//	}
	//}
	//
	//tt(data...)

	//to, err := time.Parse("2006-01-02", "2023-01-01")
	//fmt.Println(to, to.UnixNano(), err)
	//
	//data, err := base64.StdEncoding.DecodeString("UEsDBBQAAAAIACMOSlaQnA0twQAAAD4BAAA0AAAAOTY5Nzk3MjkxMV81MWMxODZiYi0wYTQ1LTQwNDMtOGI0ZS1lYzgxNTdkOTYyMGIuanNvblXPwYoCMQwG4HfJVV2atE2nua674E0Wb4uH6lQdGDsLU0ER390Oso5ev/x/Qn6v0Ie27hY1CBuNZKfQpAQCnr3zjjwiTB+Zr9SDVKw18YdST/wDMQbR+RFX4VySaF7tu0kRBCtW+FJfxhTafCkDYl/pcfATcrQgXnNl8Z2JQBwpa+zo89OwRL0H/yEcu1PKy3A5xpQfOHAdN3ket03fdMNnu9D2sWgpfh5C2pc6kCKaKZ4pv0ItmqXcdI4nwwK4re9QSwECHwAUAAAACAAjDkpWkJwNLcEAAAA+AQAANAAkAAAAAAAAACAAAAAAAAAAOTY5Nzk3MjkxMV81MWMxODZiYi0wYTQ1LTQwNDMtOGI0ZS1lYzgxNTdkOTYyMGIuanNvbgoAIAAAAAAAAQAYADcRwAELPdkBNxHAAQs92QHg9tD+Cj3ZAVBLBQYAAAAAAQABAIYAAAATAQAAAAA=")
	//if err != nil {
	//	log.Fatal("error:", err)
	//}
	//
	//ioutil.WriteFile("./filename", data, 0644)

	//var letters = []rune("abcdef0123456789")
	//rand.Seed(time.Now().UnixNano())
	//
	//b := make([]rune, 32)
	//for i := range b {
	//	b[i] = letters[rand.Intn(len(letters))]
	//
	//}
	//
	//token := string(b)
	//
	//fmt.Println(token)

	//str := []byte("MIII/jCCCKugAwIBAgIRAs3WmQClr5a3SUwU8mU6Us0wCgYIKoUDBwEBAwIwggEPMRgwFgYFKoUDZAESDTEwMjQwMDE0MzQwNDkxGjAYBggqhQMDgQMBARIMMDA0MDI5MDE3OTgxMQswCQYDVQQGEwJSVTEeMBwGA1UECAwVNDAg0JrQsNC70YPQttGB0LrQsNGPMRkwFwYDVQQHDBDQsy4g0JrQsNC70YPQs9CwMSwwKgYDVQQJDCPQv9C10YAuINCi0LXRgNC10L3QuNC90YHQutC40Lkg0LQuNjEnMCUGA1UECgwe0JDQniDQmtCQ0JvQo9CT0JAg0JDQodCi0KDQkNCbMTgwNgYDVQQDDC/QotC10YHRgtC+0LLRi9C5INCQ0J4g0JrQkNCb0KPQk9CQINCQ0KHQotCg0JDQmzAeFw0yMzAyMTAwOTEwMDdaFw0yNDA1MTAwOTIwMDdaMIICoTEuMCwGA1UECQwl0YPQuy4g0KbQuNC+0LvQutC+0LLRgdC60L7Qs9C+LNC0LiAxMDEVMBMGBSqFA2QEEgo5NjQzMTY1MjUyMR8wHQYJKoZIhvcNAQkBFhB0ZXN0QGV4YW1wbGUuY29tMRgwFgYFKoUDZAESDTExOTE1ODE3ODcwNjYxFjAUBgUqhQNkAxILMTA1MDM3ODg5MzUxGjAYBggqhQMDgQMBARIMOTY0NzI0NDc2MDEwMTwwOgYDVQQMDDPQn9C10YDQtdCy0L7RgNCw0YfQuNCy0LDRgtC10LvRjCDQv9C40L3Qs9Cy0LjQvdC+0LIxIjAgBgNVBCoMGdCi0LXRgdGCINCi0LXRgdGC0L7QstC40YcxFTATBgNVBAQMDNCi0LXRgdGC0L7QsjELMAkGA1UEBhMCUlUxLTArBgNVBAgMJDQwINCa0LDQu9GD0LbRgdC60LDRjyDQvtCx0LvQsNGB0YLRjDEZMBcGA1UEBwwQ0LMuINCa0LDQu9GD0LPQsDF2MHQGA1UECgxt0JfQsNGP0LLQutCwINC90LAg0LjQt9C80LXQvdC10L3QuNC1INCw0YLRgNC40LHRg9GC0LAg0LDQsdC+0L3QtdC90YLQsCDRgSDQv9C10YDQtdCy0YvQv9GD0YHQutC+0Lwg0YHQtdGA0YLQsDEpMCcGA1UECwwg0JDQniAi0JrQsNC70YPQs9CwINCQ0YHRgtGA0LDQuyIxdjB0BgNVBAMMbdCX0LDRj9Cy0LrQsCDQvdCwINC40LfQvNC10L3QtdC90LjQtSDQsNGC0YDQuNCx0YPRgtCwINCw0LHQvtC90LXQvdGC0LAg0YEg0L/QtdGA0LXQstGL0L/Rg9GB0LrQvtC8INGB0LXRgNGC0LAwZjAfBggqhQMHAQEBATATBgcqhQMCAiQABggqhQMHAQECAgNDAARAqAyUDHaldmdNs5RUsLFII8q6j3VWQ++wFhsbXP47B2goeHOCGDJTzPxtRppcsIJ5kkfDtjQ13466UjtWy6PCnKOCBEMwggQ/MB0GA1UdDgQWBBQwGu9MF4jjx/wS1cinNq8u8ny8+DAfBgkrBgEEAYI3FQcEEjAQBggqhQMCAi4ACAIBAQIBADAdBgNVHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwQwDAYFKoUDZHIEAwIBADAOBgNVHQ8BAf8EBAMCA/gwJwYJKwYBBAGCNxUKBBowGDAKBggrBgEFBQcDAjAKBggrBgEFBQcDBDBRBggrBgEFBQcBAQRFMEMwQQYIKwYBBQUHMAKGNWh0dHA6Ly9yZWdzZXJ2aWNlLmtleWRpc2sucnUvdGVzdGNhL3Jvb3QvdGVzdGNhMTMuY3J0MB0GA1UdIAQWMBQwCAYGKoUDZHEBMAgGBiqFA2RxAjCCAUoGBSqFA2RwBIIBPzCCATsMU9Ch0JrQl9CYICLQmtGA0LjQv9GC0L7Qn9Cg0J4gQ1NQIiAo0LLQtdGA0YHQuNGPIDQuMCkgKNC40YHQv9C+0LvQvdC10L3QuNC1IDIgLUJhc2UpDIGR0J/RgNC+0LPRgNCw0LzQvNC90L4t0LDQv9C/0LDRgNCw0YLQvdGL0Lkg0LrQvtC80L/Qu9C10LrRgSAi0KPQtNC+0YHRgtC+0LLQtdGA0Y/RjtGJ0LjQuSDRhtC10L3RgtGAICLQmtGA0LjQv9GC0L7Qn9GA0L4g0KPQpiIg0LLQtdGA0YHQuNC40LggMi4wIgwh0KHQpC8xMjQtMzM4MCDQvtGCIDExINC80LDRjyAyMDE4DC3QodCkLzEyOC0zNTkyINC+0YIgMTcg0L7QutGC0Y/QsdGA0Y8gMjAxOCDQsy4wOwYFKoUDZG8EMgww0KHQmtCX0JggItCa0YDQuNC/0YLQvtCf0YDQviIgKNCy0LXRgNGB0LjRjyA0LjApMEUGA1UdHwQ+MDwwOqA4oDaGNGh0dHA6Ly9yZWdzZXJ2aWNlLmtleWRpc2sucnUvdGVzdGNhL2NybC90ZXN0Y2ExMy5jcmwwggFRBgNVHSMEggFIMIIBRIAUY9Owge/CYLFC/MNqxXdiywXNA8uhggEXpIIBEzCCAQ8xGDAWBgUqhQNkARINMTAyNDAwMTQzNDA0OTEaMBgGCCqFAwOBAwEBEgwwMDQwMjkwMTc5ODExCzAJBgNVBAYTAlJVMR4wHAYDVQQIDBU0MCDQmtCw0LvRg9C20YHQutCw0Y8xGTAXBgNVBAcMENCzLiDQmtCw0LvRg9Cz0LAxLDAqBgNVBAkMI9C/0LXRgC4g0KLQtdGA0LXQvdC40L3RgdC60LjQuSDQtC42MScwJQYDVQQKDB7QkNCeINCa0JDQm9Cj0JPQkCDQkNCh0KLQoNCQ0JsxODA2BgNVBAMML9Ci0LXRgdGC0L7QstGL0Lkg0JDQniDQmtCQ0JvQo9CT0JAg0JDQodCi0KDQkNCbghEC1ACKAJeuNJxH//E/LK8J9jAKBggqhQMHAQEDAgNBAPBhtEIjc3A+mOwxcMLqMSGvmH0lF+2X+Yso6xwsl+EqwjwfSuBNjfXyy0KusZMmR6rW95+tze9zLiSRvvW3w08=")

	//fmt.Println(cap(str))
	//startIdx := 0
	//toIdx := 64
	//newLine := make([]byte, 0, len(str)+(len(str)/64)+1)
	//
	//for idx := range str {
	//	if idx%64 == 0 {
	//		newLine = append(newLine, str[startIdx:toIdx]...)
	//		newLine = append(newLine, []byte("\n")...)
	//		startIdx = idx + 64
	//		toIdx = idx + 64*2
	//		if len(str) < toIdx {
	//			toIdx = len(str)
	//		}
	//	}
	//}
	//
	//fmt.Println(string(newLine))

	//test(str)

}

func test(cert []byte) {
	var content []byte
	checkB64Len := 64
	var b64Cert []byte

	b64Cert = bytes.ReplaceAll(cert, []byte("\r"), nil)    // Удаляет все символы возврата кареток
	b64Cert = bytes.ReplaceAll(b64Cert, []byte("\n"), nil) // Удаляет все переносы строк

	// Добавляем переносы строк для формата
	newB64Cert := make([]byte, 0, len(b64Cert)+(len(b64Cert)/checkB64Len)+1)

	fromIdx, toIdx := 0, checkB64Len
	for idx := range b64Cert {
		if idx%checkB64Len == 0 {
			newB64Cert = append(newB64Cert, b64Cert[fromIdx:toIdx]...)
			newB64Cert = append(newB64Cert, []byte("\n")...)
			fromIdx = idx + checkB64Len
			toIdx = idx + checkB64Len*2
			if len(b64Cert) < toIdx {
				toIdx = len(b64Cert)
			}
		}
	}

	b64Cert = newB64Cert

	content = append([]byte(`-----BEGIN CERTIFICATE-----`), []byte("\n")...)
	content = append(content, b64Cert...)
	content = append(content, []byte(`-----END CERTIFICATE-----`)...)

	fmt.Println(string(content))

}
