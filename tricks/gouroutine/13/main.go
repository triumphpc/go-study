// В функции First() канал результатов не буферизован. Это значит, что возвращается только первая горутина.
//Все остальные застревают в попытке отправить свои результаты. Получается, что если у вас более одной копии (replica), то при каждом вызове происходит утечка ресурсов.

// Для решения этой проблемы можно использовать отдельный канал done

package main

//func First(query string, replicas ...Search) Result {
//	c := make(chan Result)
//	done := make(chan struct{})
//	defer close(done)
//	searchReplica := func(i int) {
//		select {
//		case c <- replicas[i](query):
//		case <-done:
//		}
//	}
//	for i := range replicas {
//		go searchReplica(i)
//	}
//
//	return <-c
//}
