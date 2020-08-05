package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type IRouter interface {
	InitRoutes()
}

type Router struct {
	mux *mux.Router
}

func (router *Router) InitRoutes() {
	router.mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode("Hello world")
	})

	http.ListenAndServe(":8000", router.mux)
}

func New() IRouter {
	return &Router{
		mux: mux.NewRouter(),
	}
}
