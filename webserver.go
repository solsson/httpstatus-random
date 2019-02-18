package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	port                      = flag.Int("port", 8080, "Port number.")
	statusCodesCommonWeighted = [...]int{
		200, 200, 200, 200, 200, 200, 200, 200, 200, 200,
		300, 301, 302, 304, 307,
		400, 401, 403, 403, 404, 404, 404, 404, 404, 410,
		500, 500, 500, 501, 503, 503, 503, 503, 503, 550}
)

func accesslog(status int, r *http.Request) {
	t := time.Now().UTC().Format("15:04:05")
	fmt.Printf("%s %d %s %s %s\n", t, status, r.RemoteAddr, r.Method, r.Header["User-Agent"])
}

func status(w http.ResponseWriter, r *http.Request) {
	status := statusCodesCommonWeighted[rand.Intn(len(statusCodesCommonWeighted))]
	w.WriteHeader(status)
	body := fmt.Sprintf("Responded randomly with status %d\n", status)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(body)))
	io.WriteString(w, body)
	accesslog(status, r)
}

func main() {
	flag.Parse()

	http.HandleFunc("/", status)

	go log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
