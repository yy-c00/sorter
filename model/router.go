package model

import (
	"github.com/labstack/echo/v4"
)

type Router interface {
	SorterRouter
	FinderRouter
}

type FinderRouter interface	{
	FindUser(echo.Context) error
}

type SorterRouter interface {
	Sort(echo.Context) error
}
