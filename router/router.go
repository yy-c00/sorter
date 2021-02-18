package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yy-c00/sorter/model"
)

//New returns an implementation of model.Router
func New() model.Router {
	return sort{NewSorterRouter(), NewAccess()}
}

type sort struct {
	model.SorterRouter
	model.FinderRouter
}

//SetMiddlewares set default middlewares
func SetMiddlewares(e *echo.Echo) {
	//CORS default config (insecure)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func SetRoutes(e *echo.Echo, r model.Router) {

	e.GET("/find", r.FindUser)
	e.POST("/sort", r.Sort)

}