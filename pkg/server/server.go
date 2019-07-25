package server

import (
	"github.com/creekorful/microgo/pkg/httputil"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler func(*http.Request) (int, interface{})

type ApiServer interface {
	PushHandler(method string, path string, handler Handler)
	Listen(address string, port int) error
}

func NewApiServer() ApiServer {
	return &gorillaApiServer{mux: mux.NewRouter()}
}

type gorillaApiServer struct {
	mux *mux.Router
}

func (gas *gorillaApiServer) PushHandler(method string, path string, handler Handler) {
	log.Println("Registering new handler: " + method + " " + path)

	gas.mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		// execute request
		status, body := handler(request)

		// override 200 OK with no content to 204 NO CONTENT
		if status == http.StatusOK && body == nil {
			status = http.StatusNoContent
		}

		_ = httputil.WriteJsonResponse(writer, status, body)
	}).Methods(method)
}

func (gas *gorillaApiServer) Listen(address string, port int) error {
	return http.ListenAndServe(address+":"+strconv.Itoa(port), gas.mux)
}
