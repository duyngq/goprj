package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct {}


var (
	muxDispatcher = mux.NewRouter().PathPrefix("/rp/api/v1/").Subrouter()
	)

func NewMuxRouter() Router {
	return &muxRouter{}
}


func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)){
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)){
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)){
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
}
func (*muxRouter) UPDATE(uri string, f func(w http.ResponseWriter, r *http.Request)){
	muxDispatcher.HandleFunc(uri, f).Methods("UPDATE")
}
func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)){
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}
func (*muxRouter) SERVE(port string){
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port,muxDispatcher)
}