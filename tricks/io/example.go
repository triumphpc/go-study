package main

import (
	"bytes"
	"io"
	"net/http"
)

func main() {
	// Send headers from bytes and read from files parallel
	r := io.MultiReader(
		bytes.NewReader([]byte("...my header...")),
		//myFile,
	)
	// Use multiply buffers
	http.Post("http://example.com", "application/octet-stream", r)
}
