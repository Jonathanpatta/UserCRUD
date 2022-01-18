package main

import (
	"net/http"
	"regexp"
	"strings"
)

type StaticHandler struct {
	EndPoint        string
	Method          string
	HandlerFunction func(http.ResponseWriter, *http.Request)
}

type DynamicHandler struct {
	DynamicPath     string
	Method          string
	HandlerFunction func(http.ResponseWriter, *http.Request, map[string]string)
}

type Router struct {
	StaticHandlers  []*StaticHandler
	DynamicHandlers []*DynamicHandler
}

func FindDynamicUrlMatch(prefix string, rawpath string) string {
	if strings.HasPrefix(rawpath, prefix) {
		data := strings.TrimPrefix(rawpath, prefix)
		data = strings.Trim(data, "/")

		return data
	}
	return ""
}

func (router *Router) RouterMiddleWare(rw http.ResponseWriter, r *http.Request) {
	for _, handler := range router.StaticHandlers {
		if handler.EndPoint == r.URL.Path && handler.Method == r.Method {
			handler.HandlerFunction(rw, r)
			return
		}
	}
	for _, handler := range router.DynamicHandlers {

		if handler.Method == r.Method {
			re := regexp.MustCompile(`(?s)\{(.*)\}`)
			dataname := re.FindString(handler.DynamicPath)
			prefix := strings.Split(handler.DynamicPath, "{")[0]
			dataname = strings.Trim(dataname, "{}")
			data := FindDynamicUrlMatch(prefix, r.URL.Path)
			if data != "" {
				datamap := make(map[string]string)
				datamap[dataname] = data
				handler.HandlerFunction(rw, r, datamap)
				return
			}

			break
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

func (r *Router) RegisterDynamicHandler(path string, method string, fn func(http.ResponseWriter, *http.Request, map[string]string)) {
	newHandler := DynamicHandler{DynamicPath: path, Method: method, HandlerFunction: fn}

	r.DynamicHandlers = append(r.DynamicHandlers, &newHandler)
}
