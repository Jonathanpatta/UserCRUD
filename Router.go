package main

import "net/http"

type Handler struct {
	EndPoint        string
	Method          string
	HandlerFunction func(http.ResponseWriter, *http.Request)
}

type Router struct {
	Handlers []*Handler
}

func (router *Router) RouterMiddleWare(rw http.ResponseWriter, r *http.Request) {
	for _, handler := range router.Handlers {
		if handler.EndPoint == r.URL.Path && handler.Method == r.Method {
			handler.HandlerFunction(rw, r)
		}
	}
}

func (r *Router) NewRouter() *Router {
	http.HandleFunc("/", r.RouterMiddleWare)
	return &Router{}
}

func (r *Router) RegisterHandler(endpoint string, method string, fn func(http.ResponseWriter, *http.Request)) {
	newHandler := Handler{EndPoint: endpoint, Method: method, HandlerFunction: fn}

	r.Handlers = append(r.Handlers, &newHandler)
}
