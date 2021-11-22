package router

import (
	"context"
	"net/http"
	"sync"

	"github.com/aws/aws-lambda-go/events"
)

var r *Router
var mu sync.Mutex

func initialize() {
	mu.Lock()
	defer mu.Unlock()
	if r == nil {
		r = New()
	}
}

func Get(pattern string, handler HandlerFunc) {
	initialize()
	r.Handle(http.MethodGet, pattern, handler)
}

func Head(pattern string, handler HandlerFunc) {
	initialize()
	r.Handle(http.MethodHead, pattern, handler)
}

func Post(pattern string, handler HandlerFunc) {
	initialize()
	r.Handle(http.MethodPost, pattern, handler)
}

func Put(pattern string, handler HandlerFunc) {
	initialize()
	r.Handle(http.MethodPut, pattern, handler)
}

func Patch(pattern string, handler HandlerFunc) {
	initialize()
	r.Handle(http.MethodPatch, pattern, handler)
}

func Delete(pattern string, handler HandlerFunc) {
	r.Handle(http.MethodDelete, pattern, handler)
}

func Options(pattern string, handler HandlerFunc) {
	initialize()
	r.Handle(http.MethodOptions, pattern, handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	initialize()
	return r.Handler(ctx, request)
}

func NewGroup(name string) *Group {
	initialize()
	g := new(Group)
	g.Name = name
	g.r = r
	return g
}
