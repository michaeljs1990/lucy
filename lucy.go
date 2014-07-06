package lucy

import (
	"net/http"
	"strings"
)

// Make everything available via service
// type Handle func(*http.Request, http.ResponseWriter, Params)

type HandlerFunc func(*Service)

type Mapper struct {
	// Hold all paths for this application
	Paths map[string][]*Service
}

// Kickstart the application and get it ready to
// start accepting routes.
func Kickstart() *Mapper {
	return &Mapper{make(map[string][]*Service)}
}

func (r *Mapper) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, method := range r.Paths[req.Method] {
		// Use .Path to get string format of URL
		if match := method.Matcher(req.URL.Path); match {
			// Write to screen
			method.R = req
			method.W = w
			method.response(method)
		}
	}
}

func (r *Mapper) Insert(method, path string, handler HandlerFunc) {
	r.Paths[method] = append(r.Paths[method], &Service{
		handler,
		path,
		Params{make(map[string]string)},
		nil,
		nil,
	})
}

// Methods below are for ease of use only
// All we do is add a layer on top of Insert

func (r *Mapper) Get(path string, handler HandlerFunc) {
	r.Insert("HEAD", path, handler)
	r.Insert("GET", path, handler)
}

func (r *Mapper) Post(path string, handler HandlerFunc) {
	r.Insert("POST", path, handler)
}

func (r *Mapper) Put(path string, handler HandlerFunc) {
	r.Insert("PUT", path, handler)
}

func (r *Mapper) Delete(path string, handler HandlerFunc) {
	r.Insert("DELETE", path, handler)
}

func (r *Mapper) Head(path string, handler HandlerFunc) {
	r.Insert("HEAD", path, handler)
}

func (r *Mapper) Options(path string, handler HandlerFunc) {
	r.Insert("OPTIONS", path, handler)
}

type Params struct {
	Params map[string]string
}

// Returned stored params
func (p *Params) Get(k string) string {
	return p.Params[k]
}

// Helper function to store params
func (p *Params) SetParams(k string, v string) {
	p.Params[k] = v
}

type Service struct {
	response HandlerFunc
	pattern  string
	Param    Params
	R        *http.Request
	W        http.ResponseWriter
}

// Write data to the screen with a response code
func (s *Service) Writer(b []byte, i int) {
	s.W.WriteHeader(i)
	s.W.Write(b)
}

// This function handles all the matching for the routing
// It will also set all the params that may come with the route
func (s *Service) Matcher(u string) bool {
	// Set match to keep track of the valid route
	var match bool = true

	is := strings.Split(u, "/")
	ps := strings.Split(s.pattern, "/")

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
			s.Param.SetParams(strings.TrimPrefix(ps[i], ":"), is[i])
		}
	}

	return match
}
