package middleware

import (
	"fmt"
	"net/http"
)

func ContentTypeFilterer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ct := req.Header.Get("Content-Type")
		if !isJSONRequest(req.Header.Get("Content-Type")) {
			http.Error(w, fmt.Sprintf("This API only allows 'application/json' requests (provided: %s).", ct), http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, req)
	})
}

func isJSONRequest(ct string) bool {
	switch ct {
	case "application/json", "application/json; charset=UTF-8":
		return true
	default:
		return false
	}
}
