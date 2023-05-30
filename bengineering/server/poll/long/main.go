package long

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faaizz/learnings/bengineering/server/common"
	"github.com/google/uuid"
)

var Requests map[string]string

type AsyncHandler struct {
	requests map[string]string
	*common.BaseAsyncHandler
}

func New() *AsyncHandler {
	return &AsyncHandler{
		requests: make(map[string]string),
	}
}

func (ah *AsyncHandler) AsyncHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ah.handlePost(w, r)
		return
	}

	ah.handleGet(w)
}

func (ah *AsyncHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	rID := r.FormValue("id")
	log.Printf("ID: %s", rID)

	status, ok := ah.requests[rID]
	if !ok {
		log.Println("ID not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if status == "" {
		log.Println("ID empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", status)
}

func (ah *AsyncHandler) handleGet(w http.ResponseWriter) {
	rID := uuid.New().String()

	ah.requests[rID] = common.Job(10, rID)

	fmt.Fprintf(w, "%s", rID)
}
