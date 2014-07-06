lucy
====

Lucy sets out to be a simplistic router written in Go that provides you everything you need in a base router.

### Example Use

Below is an example of how you can use the easy param service to get params that you have defined and pass them into your http.Handler function.

```
  func testing(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte(lucy.Param.Get("id")))
  }

  func main() {
  	temp := lucy.Diamond()
  	temp.Get("/test/:id", lucy.HandlerFunc(testing))
  	http.Handle("/", temp)
  
  	//Start Server and listen
  	log.Fatal(http.ListenAndServe(":8080", nil))
  }
```

Please take a look at the code and leave me any bugs, comments or pull requests.
