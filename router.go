package lucy

import "net/http"

//import "fmt"

type Router struct {

	//Path that should catch this route
	Paths map[string][]*Matcher
}

//Set default values for Route struct
func Diamond() *Router {
	return &Router{make(map[string][]*Matcher)}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, method := range r.Paths[req.Method] {
		// Use .Path to get string format of URL
		if match := method.Matching(req.URL.Path); match {
			t := Params{}
			t.SetTest("this")
			//Write information to screen
			method.Response.ServeHTTP(w, req)
		}
	}
}

func (r *Router) Insert(method, path string, handler http.Handler) {
	r.Paths[method] = append(r.Paths[method], &Matcher{path, handler})

	//Handle case where route may be / or nothing at all
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
	//Patern to serve up this route for
	Pattern string
	//Bool to check if this pattern should
	//check for params or not.
	//Check bool
	//Holds params that are availble for this
	//route is Check is true.
	//Params map[string]string
	//Custom Hanlder function
	Response http.Handler
}

func (m *Matcher) Matching(u string) bool {
	// if u == m.Pattern {
	// 	return true
	// } else {
	// 	return false
	// }
	return true //Only for testing purposes
}

type Params struct {
	Test string
}

func (p Params) Get(s string) string {
	return p.Test
}

func (p *Params) SetTest(s string) {
	p.Test = s
}

type HandlerFunc func(*Params) []byte

func (h HandlerFunc) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	t := &Params{}
	t.SetTest("dumpthisout")
	writer.Write(h(t))
}
