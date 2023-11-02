package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

// Хорошее решение — искусственно ограничить количество запросов на стороне сервиса.
//Для этой цели отлично подходит паттерн workerpool.
// . Предположим, надо держать исходящий RPS (requests per second) в пределах некоторого точного значения. Этого можно добиться, если построить архитектуру, где в каждый момент времени будет гарантированно выполняться не более N запросов.

type Job struct {
	URL string
}

func makeGetRequest(url string) {
	log.Println("making request to", url)

	_, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("success", url)
}

func main() {
	// возможность устанавливать число N через флаг
	numOfWorkers := flag.Int("c", 1, "number of concurrent workers")
	flag.Parse()

	jobCh := make(chan *Job)
	for i := 0; i < *numOfWorkers; i++ {
		go func() {
			for job := range jobCh {
				makeGetRequest(job.URL)
			}
		}()
	}

	// это наша DDoS-прокси
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("url")
		numOfRequestsRaw := r.FormValue("num")

		numOfRequests, err := strconv.Atoi(numOfRequestsRaw)
		if err != nil {
			http.Error(w, "wrong number", http.StatusBadRequest)
			return
		}

		for i := 0; i < numOfRequests; i++ {
			job := &Job{URL: url}
			jobCh <- job
		}

		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", nil)
}
