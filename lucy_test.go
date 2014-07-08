package lucy

import (
	"net/http"
	"testing"
)

//Create mock response writer
type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

// Test to see if the parameter service is working
// as intended and returning proper values.
func TestParams(t *testing.T) {
	p := &Params{make(map[string]string)}

	p.SetParams("key", "value")
	p.SetParams("id", "23")
	p.SetParams("name", "mike")

	if p.Get("key") != "value" {
		t.Error("SetParams did not set a key of 'key' or a value of 'value'.")
	}

	if p.Get("id") != "23" {
		t.Error("SetParams did not set a key of 'id' or a value of '23'.")
	}

	if p.Get("name") != "mike" {
		t.Error("SetParams did not set a key of 'name' or a value of 'mike'.")
	}

	if p.Get("fail") != "" {
		t.Error("SetParams unset value returned value other than empty string.")
	}
}

// Test to make sure basic routes are resolving properly
func TestRouting(t *testing.T) {
	m := Kickstart()

	routed := false

	// Test GET
	m.Get("/get", func(s *Service) {
		routed = true
	})

	//Test PUT
	m.Put("/put", func(s *Service) {
		routed = true
	})

	//Test POST
	m.Post("/post", func(s *Service) {
		routed = true
	})

	//Test DELETE
	m.Delete("/delete", func(s *Service) {
		routed = true
	})

	m.Head("/head", func(s *Service) {
		routed = true
	})

	m.Options("/options", func(s *Service) {
		routed = true
	})

	http.Handle("/", m)

	w := new(mockResponseWriter)

	req, _ := http.NewRequest("GET", "/get", nil)
	m.ServeHTTP(w, req)

	if routed == false {
		t.Error("GET route for /get has not been found.")
	}

	routed = false
	req, _ = http.NewRequest("PUT", "/put", nil)
	m.ServeHTTP(w, req)

	if routed == false {
		t.Error("PUT route for /put has not been found.")
	}

	routed = false
	req, _ = http.NewRequest("POST", "/post", nil)
	m.ServeHTTP(w, req)

	if routed == false {
		t.Error("POST route for /post has not been found.")
	}

	routed = false
	req, _ = http.NewRequest("DELETE", "/delete", nil)
	m.ServeHTTP(w, req)

	if routed == false {
		t.Error("DELETE route for /delete has not been found.")
	}

	routed = false
	req, _ = http.NewRequest("OPTIONS", "/options", nil)
	m.ServeHTTP(w, req)

	if routed == false {
		t.Error("OPTIONS route for /options has not been found.")
	}

	routed = false
	req, _ = http.NewRequest("HEAD", "/head", nil)
	m.ServeHTTP(w, req)

	if routed == false {
		t.Error("HEAD route for /head has not been found.")
	}
}

// Test that the matcher is matching the proper names to parameters
func TestMatching(t *testing.T) {
	m := Kickstart()

	matched := false

	// Test GET
	m.Get("/get/:id/:name", func(s *Service) {
		if s.Param.Get("id") == "23" && s.Param.Get("name") == "lenny" {
			matched = true
		}
	})

	//Test PUT
	m.Put("/put/:id", func(s *Service) {
		if s.Param.Get("id") == "23" {
			matched = true
		}
	})

	//Test DELETE
	m.Delete("/delete/:random", func(s *Service) {
		if s.Param.Get("random") == "zdjae" {
			matched = true
		}
	})

	//http.Handle("/", m)

	w := new(mockResponseWriter)

	req, _ := http.NewRequest("GET", "/get/23/lenny", nil)
	m.ServeHTTP(w, req)

	if matched == false {
		t.Error("GET route for /get/:id/:name has not been found or matched.")
	}

	matched = false
	req, _ = http.NewRequest("PUT", "/put/23", nil)
	m.ServeHTTP(w, req)

	if matched == false {
		t.Error("PUT route for /put/:id has not been found or matched.")
	}

	matched = false
	req, _ = http.NewRequest("DELETE", "/delete/zdjae", nil)
	m.ServeHTTP(w, req)

	if matched == false {
		t.Error("DELETE route for /delete/:random has not been found or matched.")
	}
}

// Test to ensure redirect works as intended
func TestRedirect(t *testing.T) {
	m := Kickstart()

	redirect := false

	m.Get("/original", func(s *Service) {
		s.Redirect("GET", "/redirect")
	})

	m.Get("/redirect", func(s *Service) {
		redirect = true
	})

	// Check double redirect
	m.Get("/tricky", func(s *Service) {
		s.Redirect("GET", "/tricky1")
	})

	m.Get("/tricky1", func(s *Service) {
		s.Redirect("GET", "/tricky2")
	})

	m.Get("/tricky2", func(s *Service) {
		redirect = true
	})

	w := new(mockResponseWriter)

	req, _ := http.NewRequest("GET", "/original", nil)
	m.ServeHTTP(w, req)

	if redirect == false {
		t.Error("GET /original did not redirect to GET /redirect.")
	}

	redirect = false
	req, _ = http.NewRequest("GET", "/tricky", nil)
	m.ServeHTTP(w, req)

	if redirect == false {
		t.Error("Redirect /tricky failed on double redirect.")
	}
}
