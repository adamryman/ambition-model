package svc

// This file provides server-side bindings for the HTTP transport.
// It utilizes the transport/http.Server.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	//stdopentracing "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	//"github.com/go-kit/kit/endpoint"
	//"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"

	// This service
	pb "github.com/adamryman/ambition-truss/ambition-service"
)

var (
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = strconv.Atoi
	_ = httptransport.NewServer
	_ = ioutil.NopCloser
	_ = pb.RegisterAmbitionServiceServer
	_ = io.Copy
)

// MakeHTTPHandler returns a handler that makes a set of endpoints available
// on predefined paths.
func MakeHTTPHandler(ctx context.Context, endpoints Endpoints, logger log.Logger) http.Handler {
	//func MakeHTTPHandler(ctx context.Context, endpoints Endpoints, /*tracer stdopentracing.Tracer,*/ logger log.Logger) http.Handler {
	/*options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}*/
	m := http.NewServeMux()

	return m
}

func HttpDecodeLogger(next httptransport.DecodeRequestFunc, logger log.Logger) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		logger.Log("method", r.Method, "url", r.URL.String())
		rv, err := next(ctx, r)
		if err != nil {
			logger.Log("method", r.Method, "url", r.URL.String(), "Error", err)
		}
		return rv, err
	}
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	code := http.StatusInternalServerError
	msg := err.Error()

	/*if e, ok := err.(httptransport.Error); ok {
		msg = e.Err.Error()
		switch e.Domain {
		case httptransport.DomainDecode:
			code = http.StatusBadRequest

		case httptransport.DomainDo:
			switch e.Err {
			case ErrTwoZeroes, ErrMaxSizeExceeded, ErrIntOverflow:
				code = http.StatusBadRequest
			}
		}
	}*/

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorWrapper{Error: msg})
}

func errorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}

// Server Decode

// Client Decode

// Client Encode

// EncodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeHTTPGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// PathParams takes a url and a gRPC-annotation style url template, and
// returns a map of the named parameters in the template and their values in
// the given url.
//
// PathParams does not support the entirety of the URL template syntax defined
// in third_party/googleapis/google/api/httprule.proto. Only a small subset of
// the functionality defined there is implemented here.
func PathParams(url string, urlTmpl string) (map[string]string, error) {
	rv := map[string]string{}
	pmp := BuildParamMap(urlTmpl)

	parts := strings.Split(url, "/")
	for k, v := range pmp {
		rv[k] = parts[v]
	}

	return rv, nil
}

// BuildParamMap takes a string representing a url template and returns a map
// indicating the location of each parameter within that url, where the
// location is the index as if in a slash-separated sequence of path
// components. For example, given the url template:
//
//     "/v1/{a}/{b}"
//
// The returned param map would look like:
//
//     map[string]int {
//         "a": 2,
//         "b": 3,
//     }
func BuildParamMap(urlTmpl string) map[string]int {
	rv := map[string]int{}

	parts := strings.Split(urlTmpl, "/")
	for idx, part := range parts {
		if strings.ContainsAny(part, "{}") {
			param := RemoveBraces(part)
			rv[param] = idx
		}
	}
	return rv
}

// RemoveBraces replace all curly braces in the provided string, opening and
// closing, with empty strings.
func RemoveBraces(val string) string {
	val = strings.Replace(val, "{", "", -1)
	val = strings.Replace(val, "}", "", -1)
	return val
}

// QueryParams takes query parameters in the form of url.Values, and returns a
// bare map of the string representation of each key to the string
// representation for each value. The representations of repeated query
// parameters is undefined.
func QueryParams(vals url.Values) (map[string]string, error) {

	rv := map[string]string{}
	for k, v := range vals {
		rv[k] = v[0]
	}
	return rv, nil
}
