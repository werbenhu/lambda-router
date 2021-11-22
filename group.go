package router

import (
	"net/http"
)

type Group struct {
	Name string
	r    *Router
}

func (g *Group) Get(pattern string, handler HandlerFunc) {
	pa := "/" + g.Name + pattern
	g.r.Handle(http.MethodGet, pa, handler)
}

func (g *Group) Head(pattern string, handler HandlerFunc) {
	pa := "/" + g.Name + pattern
	g.r.Handle(http.MethodHead, pa, handler)
}

func (g *Group) Post(pattern string, handler HandlerFunc) {
	pa := "/" + g.Name + pattern
	g.r.Handle(http.MethodPost, pa, handler)
}

func (g *Group) Put(pattern string, handler HandlerFunc) {
	pa := "/" + g.Name + pattern
	g.r.Handle(http.MethodPut, pa, handler)
}

func (g *Group) Patch(pattern string, handler HandlerFunc) {
	pa := "/" + g.Name + pattern
	g.r.Handle(http.MethodPatch, pa, handler)
}

func (g *Group) Delete(pattern string, handler HandlerFunc) {
	pa := "/" + g.Name + pattern
	r.Handle(http.MethodDelete, pa, handler)
}

func (g *Group) Options(pattern string, handler HandlerFunc) {
	pa := "/" + g.Name + pattern
	g.r.Handle(http.MethodOptions, pa, handler)
}
