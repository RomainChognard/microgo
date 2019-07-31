package httputil

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func CORSMethodMiddleware(r *mux.Router, allowOrigin string, allowHeaders []string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
				routeMethods, _ := route.GetMethods()

				// add OPTIONS if not already defined
				if !hasMethod(routeMethods, http.MethodOptions) {
					route.Methods(http.MethodOptions)
					routeMethods = append(routeMethods, http.MethodOptions)
				}

				w.Header().Set("Access-Control-Allow-Methods", strings.Join(routeMethods, ","))
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowHeaders, ","))
				w.Header().Set("Access-Control-Allow-Origin", allowOrigin)

				return nil
			})
			next.ServeHTTP(w, req)
		})
	}
}

func hasMethod(methods []string, method string) bool {
	for _, definedMethod := range methods {
		if definedMethod == method {
			return true
		}
	}

	return false
}
