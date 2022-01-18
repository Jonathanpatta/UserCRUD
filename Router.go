package main

import (
	"net/http"
)

type StaticHandler struct {
	EndPoint        string
	Method          string
	HandlerFunction func(http.ResponseWriter, *http.Request)
}

type DynamicHandler struct {
	DynamicPath     []string
	Method          string
	HandlerFunction func(http.ResponseWriter, *http.Request, map[string]string)
}

type Router struct {
	StaticHandlers  []*StaticHandler
	DynamicHandlers []*StaticHandler
}

func (router *Router) RouterMiddleWare(rw http.ResponseWriter, r *http.Request) {
	for _, handler := range router.StaticHandlers {
		if handler.EndPoint == r.URL.Path && handler.Method == r.Method {
			handler.HandlerFunction(rw, r)
			return
		}
	}
	http.Error(rw, "404 not found.", http.StatusNotFound)
}

func NewRouter() *Router {
	r := &Router{}
	http.HandleFunc("/", r.RouterMiddleWare)
	return r
}

func (r *Router) RegisterHandler(endpoint string, method string, fn func(http.ResponseWriter, *http.Request)) {
	newHandler := StaticHandler{EndPoint: endpoint, Method: method, HandlerFunction: fn}
	newHandlerWithSlash := StaticHandler{EndPoint: endpoint + "/", Method: method, HandlerFunction: fn}

	r.StaticHandlers = append(r.StaticHandlers, &newHandler, &newHandlerWithSlash)
}

func (r *Router) RegisterDynamicHandler(path []string, method string, fn func(http.ResponseWriter, *http.Request)) {

}
