// Sample program that performs a series of I/O related tasks to
// better understand tracing in Go.

// https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/9780135261651-UGP2_01_14_05/
// 1. ./trace                                                                                                                               triumphpc@MacBook-Pro-triumphpc
//2022/06/14 22:59:58 Searching 4000 files, found president 28000 times.

// 2. Проверяем профиль процессора
// ./trace > p.out

// 3. Смотрим трейс
/*
go tool pprof p.out                                                                                                                   triumphpc@MacBook-Pro-triumphpc
Type: cpu
Time: Jun 14, 2022 at 11:01pm (MSK)
Duration: 2.62s, Total samples = 2.51s (95.72%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)

(pprof) list freq
Total: 2.51s
ROUTINE ======================== main.freq in /Users/triumphpc/Devel/go/src/github.com/triumphpc/go-study/tricks/trace/2/trace.go
         0      2.17s (flat, cum) 86.45% of Total
         .          .     68:)
         .          .     69:
         .          .     70:type (
         .          .     71:	item struct {
         .          .     72:		XMLName     xml.Name `xml:"item"`
         .      1.86s     73:		Title       string   `xml:"title"`
         .          .     74:		Description string   `xml:"description"`
         .          .     75:	}
         .          .     76:
         .          .     77:	channel struct {
         .          .     78:		XMLName xml.Name `xml:"channel"`
         .      100ms     79:		Items   []item   `xml:"item"`
         .          .     80:	}
         .          .     81:
         .          .     82:	document struct {
         .          .     83:		XMLName xml.Name `xml:"rss"`
         .          .     84:		Channel channel  `xml:"channel"`
         .          .     85:	}
         .          .     86:)
         .      210ms     87:
         .          .     88:func main() {
         .          .     89:	pprof.StartCPUProfile(os.Stdout)
         .          .     90:	defer pprof.StopCPUProfile()
         .          .     91:
         .          .     92:	// trace.Start(os.Stdout)
(pprof) web

4. Раcскомментим trace

 go build trace.go                                                                                                                 2 ↵ triumphpc@MacBook-Pro-triumphpc
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
~/Devel/go/src/github.com/triumphpc/go-study/tricks/trace/2 (master*) » ./trace > p.out                                                                                                                       triumphpc@MacBook-Pro-triumphpc
2022/06/15 19:11:54 Searching 4000 files, found president 28000 times.

go tool trace p.out                                                                                                               2 ↵ triumphpc@MacBook-Pro-triumphpc
2022/06/15 19:13:03 Parsing trace...
2022/06/15 19:13:04 Splitting trace...
2022/06/15 19:13:04 Opening browser. Trace viewer is listening on http://127.0.0.1:58006

5. Смотрим трайс горутин:
Goroutines:
runtime.main N=1
runtime.gcBgMarkWorker N=16
runtime.bgsweep N=1
runtime.bgscavenge N=1
runtime/trace.Start.func1 N=1
N=2


*/

package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	//"runtime/pprof"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
)

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
	}

	channel struct {
		XMLName xml.Name `xml:"channel"`
		Items   []item   `xml:"item"`
	}

	document struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

func main() {
	//pprof.StartCPUProfile(os.Stdout)
	//defer pprof.StopCPUProfile()

	trace.Start(os.Stdout)
	defer trace.Stop()

	docs := make([]string, 4000)
	for i := range docs {
		docs[i] = fmt.Sprintf("newsfeed-%.4d.xml", i)
	}

	topic := "president"
	//n := freq(topic, docs)
	n := freqConcurrent(topic, docs)
	// n := freqConcurrentSem(topic, docs)
	// n := freqProcessors(topic, docs)
	// n := freqProcessorsTasks(topic, docs)
	// n := freqActor(topic, docs)

	log.Printf("Searching %d files, found %s %d times.", len(docs), topic, n)
}

func freq(topic string, docs []string) int {
	var found int

	for _, doc := range docs {
		file := fmt.Sprintf("%s.xml", doc[:8])
		f, err := os.OpenFile(file, os.O_RDONLY, 0)
		if err != nil {
			log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
			return 0
		}

		data, err := io.ReadAll(f)
		f.Close()
		if err != nil {
			log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
			return 0
		}

		var d document
		if err := xml.Unmarshal(data, &d); err != nil {
			log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
			return 0
		}

		for _, item := range d.Channel.Items {
			if strings.Contains(item.Title, topic) {
				found++
				continue
			}

			if strings.Contains(item.Description, topic) {
				found++
			}
		}
	}

	return found
}

// В этой версии мы видим, что горутины задействованы гораздо плотнее (смотрим основной трейс)
// Если детальней посмотрим на статистику выполнения горутины:
//runtime.gcStart:604
//runtime.mallocgc:1205
//runtime.growslice:272
//io.ReadAll:643
//main.freqConcurrent.func1:210

// mallocgc - говорит о том, что 1205 делаем выделение памяти в heap.
// Горутина в итоге останавливается и говорит, что идет слишком много выделения памяти, и вызывает GC
// И так каждая горутина пробует изменить heap, затем вызывается CG
// Скедулер знает, что каждая горутина делает и поэтому переключает контексты в момент выделения памяти и вызова GC

func freqConcurrent(topic string, docs []string) int {
	var found int32

	g := len(docs)
	var wg sync.WaitGroup
	wg.Add(g)

	for _, doc := range docs {
		go func(doc string) {
			// Тут выделяем локальную переменную, чтобы горутины не передавали данные через глобальный кеш и не работали с кучей
			// В данном случае счетчик работает с локальной горутино
			// Во избежании гонки, мы используем примитив синхронизации atomic
			// При использовании локальной переменной данные хранятся в L1-кеше:
			/*
				         L1    l2          L3
				G  ->C1   x
				G  ->C2   x
				G  ->C3   x
			*/
			// При окончании вычислений, происходит синхронизация с L3 для атомарного значения и вычисляется новый кэшлайан

			var lFound int32
			defer func() {
				atomic.AddInt32(&found, lFound)
				wg.Done()
			}()

			file := fmt.Sprintf("%s.xml", doc[:8])
			f, err := os.OpenFile(file, os.O_RDONLY, 0)
			if err != nil {
				log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
				return
			}

			data, err := io.ReadAll(f)
			f.Close()
			if err != nil {
				log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
				return
			}

			var d document
			if err := xml.Unmarshal(data, &d); err != nil {
				log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
				return
			}

			for _, item := range d.Channel.Items {
				if strings.Contains(item.Title, topic) {
					lFound++
					continue
				}

				if strings.Contains(item.Description, topic) {
					lFound++
				}
			}
		}(doc)
	}

	wg.Wait()
	return int(found)
}

func freqConcurrentSem(topic string, docs []string) int {
	var found int32

	g := len(docs)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan bool, runtime.GOMAXPROCS(0))

	for _, doc := range docs {
		go func(doc string) {
			ch <- true
			{
				var lFound int32
				defer func() {
					atomic.AddInt32(&found, lFound)
					wg.Done()
				}()

				file := fmt.Sprintf("%s.xml", doc[:8])
				f, err := os.OpenFile(file, os.O_RDONLY, 0)
				if err != nil {
					log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
					return
				}

				data, err := io.ReadAll(f)
				f.Close()
				if err != nil {
					log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
					return
				}

				var d document
				if err := xml.Unmarshal(data, &d); err != nil {
					log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
					return
				}

				for _, item := range d.Channel.Items {
					if strings.Contains(item.Title, topic) {
						lFound++
						continue
					}

					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
			<-ch
		}(doc)
	}

	wg.Wait()
	return int(found)
}

// Тут же мы создаем количество горутин и калалов по количеству виртуальных CPU
// Считывание документов идет через канал, буфер готорого равен количеству CPU

func freqProcessors(topic string, docs []string) int {
	var found int32

	g := runtime.GOMAXPROCS(0) // Количество горутинг по равно количеству CPU
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g) // Буфферизированный канал

	for i := 0; i < g; i++ {
		go func() {
			var lFound int32
			defer func() {
				atomic.AddInt32(&found, lFound)
				wg.Done()
			}()

			// Тут читаем из канала документы и обрабатываем
			for doc := range ch {
				file := fmt.Sprintf("%s.xml", doc[:8])
				f, err := os.OpenFile(file, os.O_RDONLY, 0)
				if err != nil {
					log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
					return
				}

				data, err := io.ReadAll(f)
				f.Close()
				if err != nil {
					log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
					return
				}

				var d document
				if err := xml.Unmarshal(data, &d); err != nil {
					log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
					return
				}

				for _, item := range d.Channel.Items {
					if strings.Contains(item.Title, topic) {
						lFound++
						continue
					}

					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
		}()
	}

	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	wg.Wait()
	return int(found)
}

func freqProcessorsTasks(topic string, docs []string) int {
	var found int32

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for i := 0; i < g; i++ {
		go func() {
			var lFound int32
			defer func() {
				atomic.AddInt32(&found, lFound)
				wg.Done()
			}()

			for doc := range ch {
				func() {
					file := fmt.Sprintf("%s.xml", doc[:8])
					ctx, tt := trace.NewTask(context.Background(), doc)
					defer tt.End()

					reg := trace.StartRegion(ctx, "OpenFile")
					f, err := os.OpenFile(file, os.O_RDONLY, 0)
					if err != nil {
						log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
						return
					}
					reg.End()

					reg = trace.StartRegion(ctx, "ReadAll")
					data, err := io.ReadAll(f)
					f.Close()
					if err != nil {
						log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
						return
					}
					reg.End()

					reg = trace.StartRegion(ctx, "Unmarshal")
					var d document
					if err := xml.Unmarshal(data, &d); err != nil {
						log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
						return
					}
					reg.End()

					reg = trace.StartRegion(ctx, "Search")
					for _, item := range d.Channel.Items {
						if strings.Contains(item.Title, topic) {
							lFound++
							continue
						}

						if strings.Contains(item.Description, topic) {
							lFound++
						}
					}
					reg.End()
				}()
			}
		}()
	}

	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	wg.Wait()
	return int(found)
}

// Тут пример, как можно разделить на разные задачи для разных горутин
func freqActor(topic string, docs []string) int {
	files := make(chan *os.File, 100)
	go func() {
		for _, doc := range docs {
			file := fmt.Sprintf("%s.xml", doc[:8])
			f, err := os.OpenFile(file, os.O_RDONLY, 0)
			if err != nil {
				log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
				break
			}
			files <- f
		}
		close(files)
	}()

	data := make(chan []byte, 100)
	go func() {
		for f := range files {
			d, err := io.ReadAll(f)
			f.Close()
			if err != nil {
				log.Printf("Reading Document [%s] : ERROR : %v", f.Name(), err)
				break
			}
			data <- d
		}
		close(data)
	}()

	rss := make(chan document, 100)
	go func() {
		for dt := range data {
			var d document
			if err := xml.Unmarshal(dt, &d); err != nil {
				log.Printf("Decoding Document : ERROR : %v", err)
				break
			}
			rss <- d
		}
		close(rss)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	var found int
	go func() {
		for d := range rss {
			for _, item := range d.Channel.Items {
				if strings.Contains(item.Title, topic) {
					found++
					continue
				}

				if strings.Contains(item.Description, topic) {
					found++
				}
			}
		}
		wg.Done()
	}()

	wg.Wait()
	return found
}
