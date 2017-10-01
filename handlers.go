package yac

import (
	"fmt"
	"net/http"
)

func notFound(w http.ResponseWriter, req *http.Request) {
	errorResponseJSON(w, http.StatusNotFound, fmt.Errorf("'%s' not found", req.URL.Path))
}

func methodNotAllowed(w http.ResponseWriter, req *http.Request) {
	errorResponseJSON(w, http.StatusMethodNotAllowed,
		fmt.Errorf("method '%s' not allowed on path '%s'", req.Method, req.URL.Path))
}

func errorResponseJSON(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": %d, "error": "%v"}`, status, err)
}
