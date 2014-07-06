lucy
====

Lucy sets out to be a simplistic router written in Go that provides you everything you need in a base router.

### Example Use

Below is an example of how you can use the easy param service to get params that you have defined and pass them into your http.Handler function.

```go

  func main() {
	t := lucy.Kickstart()

	t.Get("/test/:id", func(e *lucy.Service) {
		e.W.Write([]byte(e.Param.Get("id")))
	})

	http.Handle("/", t)

	//Start Server and listen
	log.Fatal(http.ListenAndServe(":8080", nil))
  }
```

Please take a look at the code and leave me any bugs, comments or pull requests.
