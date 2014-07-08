lucy
====
[![GoDoc](https://godoc.org/github.com/gin-gonic/gin?status.png)](http://godoc.org/github.com/michaeljs1990/lucy)
![Build Status](https://travis-ci.org/michaeljs1990/lucy.svg?branch=master)

Lucy sets out to be a simplistic router written in Go that provides you everything you need in a base router. Form more information on this package see the [GoDoc](http://godoc.org/github.com/michaeljs1990/lucy) listed here.

### Example Use

Below is an example of how you can use the easy param service to get params that you have defined and pass them into your http.Handler function.

```go

  func main() {
	t := lucy.Kickstart()

	t.Get("/test/:id", func(e *lucy.Service) {
		e.W.Write([]byte(e.Param.Get("id")))
	})

	http.Handle("/", t)

	log.Fatal(http.ListenAndServe(":8080", nil))
  }
```

Please take a look at the code and leave me any bugs, comments or pull requests.
