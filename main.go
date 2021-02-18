package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yy-c00/sorter/router"
	"github.com/yy-c00/sorter/database"
	"os"
	"log"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if database.Error() != nil {
		log.Fatal(database.Error())
	}

	e := echo.New()

	router.SetMiddlewares(e)
	router.SetRoutes(e, router.New())

	log.Fatal(e.Start(":" + port))
}