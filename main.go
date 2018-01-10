package main

import (
	"net/http"
	"time"
	"log"
)

const Timeout = time.Second

func main() {
	http.HandleFunc("/favicon.png", asset)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func asset(w http.ResponseWriter, r *http.Request) {
	log.Print("/ ASSET IN")
	defer log.Print("/ ASSET OUT")
	w.Header().Set("Content-Type", "image/png")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}
	for {
		time.Sleep(Timeout)
		_, err := w.Write([]byte{0})
		if err != nil {
			return
		}
		flusher.Flush()
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<html>
	<head>
		<title>Test</title>
		<link rel="icon" href="/favicon.png?x"/>
	</head>
</html>`))
	log.Print("/ PAGE LOAD")
}
