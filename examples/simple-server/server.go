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
