package main

import (
	"flag"
	"log"

	"github.com/faaizz/learnings/bengineering/client/block"
	"github.com/faaizz/learnings/bengineering/client/noblock"
)

var mode string
var err error

func init() {
	flag.StringVar(&mode, "mode", "block", "client type")
}

func main() {
	bUrl := "http://localhost:8080"
	jobPath := "/job"

	flag.Parse()
	log.Printf("mode: %s", mode)

	switch mode {
	default:
		log.Fatal("please provide a valid mode")
	case "block":
		err = block.Request(bUrl, jobPath, 5, 10)
	case "noblock":
		err = noblock.Request(bUrl, jobPath)
	}

	if err != nil {
		log.Fatal(err)
	}

}
