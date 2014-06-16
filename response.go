package lucy

import "net/http"

// Layer ontop of http.ResponseWriter
type ResponseWriter interface {

	//Include https implimentation
	http.ResponseWriter

	//Get proper information from Gem
	Params(string) string
}

// The following data structure is used to hold
// params that are passed in through the url so
// the user can access them more readily

//Impliment the lucy.ResponseWriter Interface
type LucyWriter type {

	//Impliment Generic Response Writter Functions
	http.ResponseWriter

	//The gem will hold all Params available
	//url strings that you may need to access
	Params Gem
}

type Gem struct {
	Input map[string]string
}

func (rw *ResponseWriter) Params(s string) string {
	return "hello"
}
