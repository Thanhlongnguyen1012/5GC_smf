package context

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Request struct {
	Params map[string]string
	Header http.Header
	Body   interface{}
	Query  url.Values
	URL    *url.URL
}

func NewRequest(req *http.Request, body interface{}) *Request {
	ret := &Request{}
	ret.Header = req.Header
	ret.Body = body
	ret.Query = req.URL.Query()
	ret.Params = make(map[string]string)
	ret.URL = req.URL
	return ret
}

type Response struct {
	Header http.Header
	Status int
	Body   interface{}
}

func NewResponse(h http.Header, body interface{}, code int) *Response {
	ret := &Response{}
	ret.Header = h
	ret.Body = body
	return ret
}
func NewHttp2Server(bindAddr string, handler http.Handler) *http.Server {
	if handler == nil {
		fmt.Println("server need handler to handler")
	}
	h2Server := &http2.Server{
		// TODO: extends the idle time after re-use openapi client
		IdleTimeout: 1 * time.Millisecond,
	}
	server := &http.Server{
		Addr:    bindAddr,
		Handler: h2c.NewHandler(handler, h2Server),
	}
	return server
}
