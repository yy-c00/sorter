package router

import (
	"github.com/labstack/echo/v4"
	"github.com/yy-c00/sorter/model"
	"github.com/yy-c00/sorter/sorter"
	"net/http"
)

//NewSorterRouter returns an implementation of model.SorterRouter
func NewSorterRouter() model.SorterRouter {
	return quickSort{sorter.New()}
}

type quickSort struct{
	model.Sorter
}

func (s quickSort) Sort(c echo.Context) error {
	numbers := model.Numbers{}

	err := c.Bind(&numbers)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	s.Sorter.Sort(&numbers)

	return c.JSON(http.StatusOK, numbers)
}