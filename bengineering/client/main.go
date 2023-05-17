package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/faaizz/learnings/bengineering/client/async"
)

var mode string
var handlerFunc func(http.ResponseWriter, *http.Request)

func init() {
	flag.StringVar(&mode, "mode", "async", "client type")
}

func main() {
	flag.Parse()
	log.Printf("mode: %s", mode)

	switch mode {
	default:
		log.Fatal("please provide a valid mode")
	case "async":
		async.Request()
	}
}
