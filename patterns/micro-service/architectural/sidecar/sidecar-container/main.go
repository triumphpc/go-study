package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	fmt.Fprintf(w, "Request logged")

	proxyHandler(w, r)
}

func main() {
	http.HandleFunc("/", logHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Proxying request: %s %s", r.Method, r.URL.Path)

	// Перенаправляем запрос на основное приложение
	resp, err := http.Get("http://localhost:8080" + r.URL.Path)
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Копируем ответ от основного приложения обратно клиенту
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
