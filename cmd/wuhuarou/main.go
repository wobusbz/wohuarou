package main

import (
	"github.com/wobusbz/wohuarou/http"
	"log"
	http2 "net/http"
)

func main() {
	http.RegisterRoutes()

	log.Fatal(http2.ListenAndServe("0.0.0.0:8899", nil))
}
