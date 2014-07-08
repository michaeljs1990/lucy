package lucy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Make everything available via service
// type Handle func(*http.Request, http.ResponseWriter, Params)

type HandlerFunc func(*Service)

type JS map[string]interface{}

type Mapper struct {
	// Hold all paths for this application
	Paths map[string][]*Service
}

// make mapping available for redirects and to ensure we
// never have more than one mapper in existance.
var mappings *Mapper

// Kickstart the application and get it ready to
// start accepting routes. This uses the singleton
// design pattern to ensure only one mapper exists.
func Kickstart() *Mapper {
	if mappings == nil {
		mappings = &Mapper{make(map[string][]*Service)}
		return mappings
	} else {
		return mappings
	}
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

func (r *Mapper) insert(method, path string, handler HandlerFunc) {
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
	r.insert("HEAD", path, handler)
	r.insert("GET", path, handler)
}

func (r *Mapper) Post(path string, handler HandlerFunc) {
	r.insert("POST", path, handler)
}

func (r *Mapper) Put(path string, handler HandlerFunc) {
	r.insert("PUT", path, handler)
}

func (r *Mapper) Delete(path string, handler HandlerFunc) {
	r.insert("DELETE", path, handler)
}

func (r *Mapper) Head(path string, handler HandlerFunc) {
	r.insert("HEAD", path, handler)
}

func (r *Mapper) Options(path string, handler HandlerFunc) {
	r.insert("OPTIONS", path, handler)
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

// Redirect to another route to finish
// the rest of the request.
func (s *Service) Redirect(method, path string) {
	// Change to the correct method and to the
	// correct path. May need to update a few more
	// variables down the line.
	s.R.Method = method
	s.R.URL.Path = path

	for _, method := range mappings.Paths[method] {
		// Use .Path to get string format of URL
		if match := method.Matcher(path); match {
			// Write to screen
			method.R = s.R
			method.W = s.W
			method.response(method)
		}
	}
}

// Use abort if you intend for the code running it to be
// the terminating statement in the function such as when
// throwing a 404 or 500 error.
func (s *Service) Abort(code int) {
	s.W.WriteHeader(code)
}

// Kill the process and output dump to the screen.
// Useful for trouble shooting your application.
func (s *Service) Kill(code int, dump interface{}) {
	s.W.Write([]byte(fmt.Sprint("%v", dump)))
	s.Abort(code)
}

// Write JSON to screen for the user.
func (s *Service) JSON(code int, output JS) {
	// Set Proper Header
	s.W.Header().Set("Content-Type", "application/json")
	s.W.WriteHeader(code)

	encoder := json.NewEncoder(s.W)
	if err := encoder.Encode(output); err != nil {
		s.Abort(500)
	}
}
