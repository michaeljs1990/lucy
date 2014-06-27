package lucy

import (
	"net/http"
	"strings"
)

type Router struct {
	// Path that should catch this route
	Paths map[string][]*Matcher
}

// Set default values for Route struct
func Diamond() *Router {
	return &Router{make(map[string][]*Matcher)}
}

// No heavy lifting we will leave that for method to take care of
// Here we loop through the current method to see if it matches
// any routes that have been set.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, method := range r.Paths[req.Method] {
		// Use .Path to get string format of URL
		if match := method.Matching(req.URL.Path); match {
			// Write information to screen
			method.Response.ServeHTTP(w, req)
		}
	}
}

func (r *Router) Insert(method, path string, handler http.Handler) {
	r.Paths[method] = append(r.Paths[method], &Matcher{path, handler})
	// TODO Handle case where route may be / or nothing at all
}

// Methods below are for ease of use only
// All we do is add a layer on top of Insert

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
	// Patern to serve up this route for
	Pattern string
	// Custom Handler function
	Response http.Handler
}

// Want to improve this however i can't think of
// a nice way to do it right now that would be more
// efficent. I am sure many exist though.
func (m *Matcher) Matching(u string) bool {
	// Set match to keep track of the valid route
	var match bool = true

	is := strings.Split(u, "/")
	ps := strings.Split(m.Pattern, "/")

	// Determine if this pattern is one that matches or not.
	// End loop if pattern does not match.
	for i := 0; i < len(is); i++ {
		// Check if index will be out of bounds
		// Check if arrays are the same size
		if len(ps) == i || len(is) != len(ps) {
			match = false
			break
		}

		// Check if this is a route param
		index := strings.Index(ps[i], ":")

		// Test for matching route
		if is[i] != ps[i] && index != 0 {
			match = false
			break
		}

		// Set params for use later tada!
		if index == 0 {
			Param.SetParams(strings.TrimPrefix(ps[i], ":"), is[i])
		}
	}

	return match

}

// Make Param struct available to entire namespace
var Param Params = Params{make(map[string]string)}

// Param service that will hold all possible params
// that the user has defined and make them available
// to grab with Params.Get()
type Params struct {
	Params map[string]string
}

// Returned stored params
func (p Params) Get(k string) string {
	return p.Params[k]
}

// Helper function to store params
func (p *Params) SetParams(k string, v string) {
	p.Params[k] = v
}

// Custom handler func to allow for param injection
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//Custom implimentation incase needed later.
func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}
