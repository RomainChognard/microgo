# microgo

Little go framework to help building REST api

## examples

### using ApiServer

```go
package main

import (
	"github.com/creekorful/microgo/pkg/server"
	"log"
	"net/http"
)

type UserDto struct {
	Username string
	Password string
}

func main() {
	// create instance of api server
	s := server.NewApiServer()

	// register endpoints
	s.PushHandler(http.MethodGet, "/users", getUser)
	s.PushHandler(http.MethodPost, "/users", createUser)

	// finally listen
	if err := s.Listen("0.0.0.0", 7777); err != nil {
		log.Println("Unable to listen: ", err)
	}
}

func getUser(r *http.Request) (int, interface{}) {
	return http.StatusOK, UserDto{Username: "creekorful", Password: "test"}
}

func createUser(_ *http.Request) (int, interface{}) {
	return http.StatusOK, nil
}
```

### using gorilla/mux

```go
package main

import (
	"github.com/creekorful/microgo/pkg/httputil"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UserDto struct {
	Username string
	Password string
}

func main() {
	// create instance of gorilla/mux
	r := mux.NewRouter()

	// register endpoints
	r.HandleFunc("/users", httputil.JsonHandler(getUser)).Methods(http.MethodGet)
	r.HandleFunc("/users", httputil.JsonHandler(createUser)).Methods(http.MethodPost)

	// finally listen
	if err := http.ListenAndServe("0.0.0.0:7777", r); err != nil {
		log.Println("Unable to listen: ", err)
	}
}

func getUser(r *http.Request) (int, interface{}) {
	return http.StatusOK, UserDto{Username: "creekorful", Password: "test"}
}

func createUser(_ *http.Request) (int, interface{}) {
	return http.StatusOK, nil
}

```