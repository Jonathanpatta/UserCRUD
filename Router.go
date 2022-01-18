package main

import "net/http"

type Handler struct {
	EndPoint string
	Method   string
	handler  func(http.ResponseWriter, *http.Request)
}

type Router struct {
	Handlers []*Handler
}

func (r *Router) NewRouter() *Router {
	return &Router{}
}
