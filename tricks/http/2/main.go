// Некоторые HTTP-серверы какое-то время держат сетевые соединения открытыми (согласно спецификации HTTP 1.1 и серверной конфигурации keep alive).
//По умолчанию стандартная HTTP-библиотека закрывает соединения, только когда об этом просит целевой HTTP-сервер.
//Тогда при определённых условиях в вашем приложении могут закончиться сокеты / файловые дескрипторы.
//
//Можно попросить библиотеку закрывать соединение после завершения вашего запроса, задав значение true в поле Close переменной запроса.
// Можно ещё глобально отключить повторное использование HTTP-соединений. Для этого создайте
//кастомную конфигурацию HTTP-транспорта.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("http://golang.org")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(string(body)))
}
