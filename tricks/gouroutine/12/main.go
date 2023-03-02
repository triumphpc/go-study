// В функции First() канал результатов не буферизован. Это значит, что возвращается только первая горутина.
//Все остальные застревают в попытке отправить свои результаты. Получается, что если у вас более одной копии (replica),
//то при каждом вызове происходит утечка ресурсов.

// Для решения этой проблемы можно использовать отдельный канал done
// make(chan struct{}) - используется для метки финиша, потому что struct не выделяет память
package main

import "fmt"

type Search func(str string) Result

type Result string

func main() {
	replics := make([]Search, 2)
	fn := func(str string) Result {
		return Result(str)
	}

	replics[0] = fn
	replics[1] = fn

	res := First("test", replics...)
	fmt.Println(res)

}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	done := make(chan struct{}) // Метка для того, что выходим
	defer close(done)

	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		case <-done:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}
