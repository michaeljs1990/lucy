lucy
====

Lucy sets out to be a simplistic router in go that provides you everything you need in a base router.

###Example Use

Below is an example of how you can use the easy param service to get params that you have defined and pass them into your http.Handler function.

```
  package main
  
  import (
  	"log"
  	"net/http"
  	"github.com/michaeljs1990/lucy"
  )
  
  func vans(s lucy.Params) ([]byte, int) {
  	return []byte(s.Get("id")), 404
  }
  
  func main() {
  	temp := lucy.Diamond()
  	temp.Get("/vans/:id", lucy.HandlerFunc(vans))
  	http.Handle("/", temp)
  
  	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Please take a look at the code and leave me any bugs, comments or pull requests.
