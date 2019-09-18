package main

import "net/http"

type middleware func(handler http.Handler) http.Handler
type Router struct {
	middlewareChain []middleware
	mux map[string]http.Handler
}

func NewRouter() *Router  {
	return &Router{}
}
func (r *Router) Use(m middleware)  {
	r.middlewareChain = append(r.middlewareChain, m)
}
func (r *Router) Add(route string, h http.Handler)  {
	var mergedHandler =  h

	for i := len(r.middlewareChain) - 1; i>= 0 ; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}
	r.mux[route] = mergedHandler
}

/*func main() {
	r := NewRouter()
	r.Use(logger)
	r.Use(timeout)
	r.Use(ratelimit)
	r.Add("/", helloHandler)
}*/