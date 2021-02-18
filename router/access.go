package router

import (
	"github.com/labstack/echo/v4"
	"github.com/yy-c00/sorter/database"
	"github.com/yy-c00/sorter/model"
	"net/http"
	"strconv"
)

//NewAccess returns an implementation of model.FinderRouter
func NewAccess() model.FinderRouter {
	return access{database.NewUserFinder()}
}

type access struct{
	model.UserFinder
}

func (a access) FindUser(c echo.Context) error {
	user, err := model.User{}, error(nil)

	user.Id, err = strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "missing valid query param 'id'"})
	}

	err = a.UserFinder.FindUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
