package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/faaizz/learnings/bengineering/server/async"
)

var mode string
var handlerFunc func(http.ResponseWriter, *http.Request)

func init() {
	flag.StringVar(&mode, "mode", "async", "server type")
}

func main() {
	flag.Parse()
	log.Printf("mode: %s", mode)

	switch mode {
	default:
		log.Fatal("please provide a valid mode")
	case "async":
		handlerFunc = async.Handler
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})
	http.HandleFunc("/job", func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
