package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/faaizz/learnings/bengineering/server/poll/long"
	"github.com/faaizz/learnings/bengineering/server/poll/short"
)

var mode string
var ah AsyncHandler

type AsyncHandler interface {
	RootHandle(w http.ResponseWriter, r *http.Request)
	AsyncHandle(w http.ResponseWriter, r *http.Request)
}

func init() {
	flag.StringVar(&mode, "mode", "short", "server type")
}

func main() {
	flag.Parse()
	log.Printf("mode: %s", mode)

	switch mode {
	default:
		log.Fatal("please provide a valid mode")
	case "short":
		ah = short.New()
	case "long":
		ah = long.New()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ah.RootHandle(w, r)
	})
	http.HandleFunc("/job", func(w http.ResponseWriter, r *http.Request) {
		ah.AsyncHandle(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
