package lucy

import "net/http"

type Router struct {
	//Path that should catch this route
	paths map[string][]*Matcher
}

//Set default values for Route struct
func Diamond() *Router {
	return &Router{make(map[string][]*Matcher)}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, method := range r.paths[req.Method] {
		// Use .Path to get string format of URL
		if req.URL.Path == method.pattern {
			method.ServeHTTP(w, req)
		}
	}
}

func (r *Router) Insert(method, path string, handler http.Handler) {
	r.paths[method] = append(r.paths[method], &Matcher{path, handler})

	n := len(path)
	if n > 0 && path[n-1] == '/' {
		r.Insert(method, path[:n-1], http.RedirectHandler(path, http.StatusMovedPermanently))
	}
}

// Methods bellow are for ease of use only
// All we do is add a layer ontop of Insert

func (r *Router) Get(path string, handler http.Handler) {
	r.Insert("HEAD", path, handler)
	r.Insert("GET", path, handler)
}

func (r *Router) Post(path string, handler http.Handler) {
	r.Insert("POST", path, handler)
}

func (r *Router) Put(path string, handler http.Handler) {
	r.Insert("PUT", path, handler)
}

func (r *Router) Delete(path string, handler http.Handler) {
	r.Insert("DELETE", path, handler)
}

func (r *Router) Head(path string, handler http.Handler) {
	r.Insert("HEAD", path, handler)
}

func (r *Router) Options(path string, handler http.Handler) {
	r.Insert("OPTIONS", path, handler)
}

// Matcher holds information for each route added
// to the router. http.Handle provides us with the
// ability to call ServeHTTP on any of these stucts.

type Matcher struct {
	pattern string
	http.Handler
}
