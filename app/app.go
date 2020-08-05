package app

import "github.com/bubeha/go-blog-api/app/router"

type IApplication interface {
	Init()
}

type Application struct {
	router router.IRouter
}

func (application *Application) Init() {
	application.router.InitRoutes()
}

func New() IApplication {
	return &Application{
		router: router.New(),
	}
}
