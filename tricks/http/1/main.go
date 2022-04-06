// Делая запрос с помощью стандартной HTTP-библиотеки, вы получаете переменную HTTP-ответа. Даже если вы не читаете
//тело ответа, всё равно нужно его закрыть. Обратите внимание: это относится и к пустым ответам

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
