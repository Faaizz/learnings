package common

import (
	"log"
	"time"
)

func Job(w int32, rID string) string {
	d := time.Duration(w) * time.Second
	time.Sleep(d)
	log.Printf("%s done", rID)
	return "done"
}
