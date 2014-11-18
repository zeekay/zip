package zip

import (
	"net/http"
)

type Handler func(Req, Res)
type Router func(http.ResponseWriter, *http.Request)

var (
	handlers = map[string]map[string]Handler{
		"":        make(map[string]Handler),
		"HEAD":    make(map[string]Handler),
		"GET":     make(map[string]Handler),
		"POST":    make(map[string]Handler),
		"DELETE":  make(map[string]Handler),
		"PUT":     make(map[string]Handler),
		"PATCH":   make(map[string]Handler),
		"OPTIONS": make(map[string]Handler),
	}
	routers = make(map[string]Router)
)

func AddHandler(pattern, method string, handler Handler) {
	// Add handler for this method/pattern
	handlers[method][pattern] = handler

	// Return if we've already attached a handlerFunc for this pattern.
	if _, ok := routers[pattern]; ok {
		return
	}

	router := (func(w http.ResponseWriter, r *http.Request) {
		// Create request
		req := Req{r}

		// Create response
		header := w.Header()
		res := Res{w, w, header}

		// Set header
		header.Set("Server", "zip.go")

		// handlers added with Route() get all requests
		if method == "" {
			handler(req, res)
			return
		}

		// method handlers bail if method doesn't match
		if handler, ok := handlers[r.Method][pattern]; ok {
			handler(req, res)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Add router
	routers[pattern] = router

	// register router for this pattern
	http.HandleFunc(pattern, router)
}

func Route(pattern string, handler Handler) {
	AddHandler(pattern, "", handler)
}

func Head(pattern string, handler Handler) {
	AddHandler(pattern, "HEAD", handler)
}

func Get(pattern string, handler Handler) {
	AddHandler(pattern, "GET", handler)
}

func Post(pattern string, handler Handler) {
	AddHandler(pattern, "POST", handler)
}

func Delete(pattern string, handler Handler) {
	AddHandler(pattern, "DELETE", handler)
}

func Put(pattern string, handler Handler) {
	AddHandler(pattern, "PUT", handler)
}

func Patch(pattern string, handler Handler) {
	AddHandler(pattern, "PATCH", handler)
}

func Options(pattern string, handler Handler) {
	AddHandler(pattern, "OPTIONS", handler)
}
