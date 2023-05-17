package async

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handlePost(w, r)
		return
	}

	handleGet(w)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	rID := r.FormValue("id")
	log.Printf("ID: %s", rID)

	status, ok := Requests[rID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if status == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s\n", status)
}

func handleGet(w http.ResponseWriter) {
	rID := uuid.New().String()

	Requests[rID] = "pending"
	go func() {
		Requests[rID] = job(120)
	}()

	fmt.Fprintf(w, "%s\n", rID)
}
