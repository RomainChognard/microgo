package main

import (
	"github.com/RomainChognard/microgo/pkg/httputil"
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
	r.HandleFunc("/users", httputil.JsonHandler(getUser)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/users", httputil.JsonHandler(createUser)).Methods(http.MethodPost, http.MethodOptions)

	// add CORS middleware
	r.Use(httputil.CORSMiddleware(r, "*", []string{"*"}))

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
