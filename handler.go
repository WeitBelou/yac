package yac

import "net/http"

type Handler http.HandlerFunc

var emptyHandler = Handler(func(_ http.ResponseWriter, _ *http.Request) {})
