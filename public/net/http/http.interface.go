package http

import "net/http"

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

// ResponseWriter ...
type ResponseWriter interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher

	// Returns the HTTP response status code of the current request.
	Status() int

	// Returns the number of bytes already written into the response http body.
	// See Written()
	Size() int

	// Writes the string into the response body.
	WriteString(string) (int, error)

	// Returns true if the response body was already written.
	Written() bool

	// Forces to write the http header (status code + headers).
	WriteHeaderNow()

	// get the http.Pusher for server push
	Pusher() http.Pusher
}

// ErrorType is an unsigned 64-bit error code as defined in the gin spec.
type ErrorType uint

// Error represents a error's specification.
type Error struct {
	Err  error
	Type ErrorType
	Meta interface{}
}

type errorMsgs []*Error

type Context struct {
	Request  *http.Request          // Current http request
	Params   Params                 //  a list of named path parameters.
	Writer   ResponseWriter         // the Response that we are writing to.
	Keys     map[string]interface{} // a key/value pair exclusively for the context of each request.
	Errors   errorMsgs              // Errors is a list of errors attached to all the handlers/middlewares who used this context.
	Accepted []string               // Accepted defines a list of manually accepted formats for content negotiation.
}

type Engine struct {
}

type IHTTP interface {
	New() *Engine // returns a new Engine instance
}
