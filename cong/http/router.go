package router

import "net/http"

type Router interface {
	GET(url string, f func(w http.ResponseWriter, r *http.Request))
	POST(url string, f func(w http.ResponseWriter, r *http.Request))
	PUT(url string, f func(w http.ResponseWriter, r *http.Request))
	UPDATE(url string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(url string, f func(w http.ResponseWriter, r *http.Request))

	SERVE(port string)
}
