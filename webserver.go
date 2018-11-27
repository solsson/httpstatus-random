package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

var (
	port                      = flag.Int("port", 8080, "Port number.")
	statusCodesCommonWeighted = [...]int{
		200, 200, 200, 200, 200, 200, 200, 200, 200, 200,
		300, 301, 302, 304, 307,
		400, 401, 403, 403, 404, 404, 404, 404, 404, 410,
		500, 500, 500, 501, 503, 503, 503, 503, 503, 550}
)

func status(w http.ResponseWriter, r *http.Request) {
	status := statusCodesCommonWeighted[rand.Intn(len(statusCodesCommonWeighted))]
	w.WriteHeader(status)
	io.WriteString(w, fmt.Sprintf("Responded randomly with status %d\n", status))
}

func main() {
	flag.Parse()

	http.HandleFunc("/", status)

	go log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
