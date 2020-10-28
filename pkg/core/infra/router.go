package infra

import "net/http"

type Router interface {
	AddRoute(method string, pattern string, handler http.Handler)
}
