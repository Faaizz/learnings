package common

import (
	"fmt"
	"net/http"
	"strings"
)

type BaseAsyncHandler struct{}

var acceptedHeaders []string
var acceptedMethods []string

func init() {
	acceptedHeaders = []string{
		"Content-Type",
	}

	acceptedMethods = []string{
		"GET",
		"POST",
		"OPTIONS",
	}
}

func (bah *BaseAsyncHandler) RootHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
		w.Header().Set(
			"Access-Control-Allow-Headers",
			strings.Join(acceptedHeaders, ", "),
		)

		w.Header().Set(
			"Access-Control-Allow-Methods",
			strings.Join(acceptedMethods, ", "),
		)

		w.WriteHeader(http.StatusNoContent)
	}

	fmt.Fprint(w, "Hello!")
}
