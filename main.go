package main

import (
	"log"
	"net/http"
	"time"
)

const Timeout = time.Second

func main() {
	http.HandleFunc("/favicon.ico", asset)
	http.HandleFunc("/page1", page1)
	http.HandleFunc("/page2", page2)
	http.HandleFunc("/", index)
	log.Println("Serving on :8080")
	http.ListenAndServe(":8080", nil)
}

func asset(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	t0 := time.Now()
	query := r.Header.Get("referer")
	log.Println("IN", query)
	defer func() {
		t1 := time.Now()
		log.Println("OUT", query, t1.Sub(t0))
	}()
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
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

func page1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<html>
	<head>
		<title>Page 1</title>
	</head>
	<body>
		<a href="/">Index</a> - <a href="/page1">Page 1</a> - <a href="/page2">Page 2</a>
	</body>
</html>`))
}

func page2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<html>
	<head>
		<title>Page 2</title>
	</head>
	<body>
		<a href="/">Index</a> - <a href="/page1">Page 1</a> - <a href="/page2">Page 2</a>
	</body>
</html>`))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<html>
	<head>
		<title>Index</title>
	</head>
	<body>
		<a href="/">Index</a> - <a href="/page1">Page 1</a> - <a href="/page2">Page 2</a>
	</body>
</html>`))
}
