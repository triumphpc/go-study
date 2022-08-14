package main

import (
	"encoding/json"
	"fmt"
)

type EDOFileImport struct {
	AbonentID string      `json:"ИдАбонента"`
	SignDss   *EDOSignDss `json:"ПараметрыПодписиDSS,omitempty"`
}

type EDOSignDss struct {
	CertID     string `json:"ИдСертификата"`
	WithNotify *bool  `json:"ОтправитьУведомление,omitempty"`
}

func main() {
	arg := new(EDOFileImport)
	data := `{"ИдАбонента":"3", "ПараметрыПодписиDSS":{"ИдСертификата": "444", "ОтправитьУведомление": true}}`

	if err := json.Unmarshal([]byte(data), arg); err != nil {
	}

	if arg.SignDss != nil {
		fmt.Println(arg.SignDss.CertID)
		if arg.SignDss.WithNotify != nil {
			fmt.Println(*arg.SignDss.WithNotify)

		}
	}

	fmt.Printf("%", arg)

	tt := false

	fmt.Printf("%t", tt)

	ff := []byte("ccccccc")
	cccc := []byte("gggggg")
	var content []byte

	content = append(ff, []byte("\n")...)
	content = append(content, cccc...)

	fmt.Println(string(content))

}
