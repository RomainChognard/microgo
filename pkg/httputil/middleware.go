package httputil

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func CORSMiddleware(r *mux.Router, allowOrigin string, allowHeaders []string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			var routeMethods []string
			_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
				if route.Match(req, &mux.RouteMatch{}) {
					currentMethods, _ := route.GetMethods()
					routeMethods = append(routeMethods, currentMethods...)
				}

				w.Header().Set("Access-Control-Allow-Methods", strings.Join(removeDuplicates(routeMethods), ","))
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowHeaders, ","))
				w.Header().Set("Access-Control-Allow-Origin", allowOrigin)

				return nil
			})
			next.ServeHTTP(w, req)
		})
	}
}

func removeDuplicates(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
		if _, ok := m[item]; !ok {
			m[item] = true
		}
	}
	var result []string
	for item := range m {
		result = append(result, item)
	}
	return result
}
