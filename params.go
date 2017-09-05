package yac

import (
	"context"
	"errors"
	"net/http"
)

// Wrapper around query parameters, request body and custom path params
type params struct {
	Query      map[string]string
	Body       []byte
	PathParams map[string]interface{}
}

// Extracts "params" from request
func Params(req *http.Request) (params, error) {
	if p, ok := req.Context().Value("params").(params); ok {
		return p, nil
	}
	return params{}, errors.New("can not extract params from request's context")
}

// Puts params in request context and returns new request
func putParamsToRequest(req *http.Request, params params) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), "params", params))
}
