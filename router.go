package mux

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type HandlerFunc func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// Mux is a tire base HTTP request router which can be used to
// dispatch requests to different handler functions.
type Router struct {
	trie           *Trie
	defaultHandler HandlerFunc
}

// New returns a Mux instance.
func New(opts ...Options) *Router {
	return &Router{trie: NewTrie(opts...)}
}

// Get registers a new GET route for a path with matching handler in the Mux.
func (r *Router) Get(pattern string, handler HandlerFunc) {
	r.Handle(http.MethodGet, pattern, handler)
}

// Head registers a new HEAD route for a path with matching handler in the Mux.
func (r *Router) Head(pattern string, handler HandlerFunc) {
	r.Handle(http.MethodHead, pattern, handler)
}

// Post registers a new POST route for a path with matching handler in the Mux.
func (r *Router) Post(pattern string, handler HandlerFunc) {
	r.Handle(http.MethodPost, pattern, handler)
}

// Put registers a new PUT route for a path with matching handler in the Mux.
func (r *Router) Put(pattern string, handler HandlerFunc) {
	r.Handle(http.MethodPut, pattern, handler)
}

// Patch registers a new PATCH route for a path with matching handler in the Mux.
func (r *Router) Patch(pattern string, handler HandlerFunc) {
	r.Handle(http.MethodPatch, pattern, handler)
}

// Delete registers a new DELETE route for a path with matching handler in the Mux.
func (r *Router) Delete(pattern string, handler HandlerFunc) {
	r.Handle(http.MethodDelete, pattern, handler)
}

// Options registers a new OPTIONS route for a path with matching handler in the Mux.
func (r *Router) Options(pattern string, handler HandlerFunc) {
	r.Handle(http.MethodOptions, pattern, handler)
}

// DefaultHandler registers a new handler in the Mux
// that will run if there is no other handler matching.
func (r *Router) DefaultHandler(handler HandlerFunc) {
	r.defaultHandler = handler
}

// Handle registers a new handler with method and path in the Mux.
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
func (r *Router) Handle(method, pattern string, handler HandlerFunc) {
	if method == "" {
		panic(fmt.Errorf("invalid method"))
	}
	r.trie.Parse(pattern).Handle(strings.ToUpper(method), handler)
}

// ServeHTTP implemented http.Handler interface
func (r *Router) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var handler HandlerFunc
	path := request.Path
	method := request.HTTPMethod
	match, err := r.trie.Match(path)
	if err != nil {
		return HandleError(RouterError{
			Code:    http.StatusNotFound,
			Message: "No such resource",
		})
	}
	if match.Node == nil {
		// Redirect for slash url
		// Router /a/b   Access PATH /a/b/ Redirect to /a/b
		// Router /a/b/  Access PATH /a/b  Redirect to /a/b/
		if match.Path != "" {
			request.Path = match.Path
			return r.Handler(ctx, request)
		}
		if r.defaultHandler == nil {
			return HandleError(RouterError{
				Code:    http.StatusNotFound,
				Message: "No such resource",
			})
		}
		handler = r.defaultHandler
	} else {
		var ok bool
		if handler, ok = match.Node.GetHandler(method).(HandlerFunc); !ok {
			return HandleError(RouterError{
				Code:    http.StatusNotFound,
				Message: "No such resource",
			})
		}
	}
	if match.Params != nil {
		if request.QueryStringParameters == nil {
			request.QueryStringParameters = make(map[string]string)
		}

		for k, v := range match.Params {
			request.QueryStringParameters[k] = v
		}
	}
	return handler(ctx, request)
}
