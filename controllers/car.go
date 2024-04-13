package controllers

import (
	"github.com/kimcodec/test_api_task/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CarService interface {
	List(c echo.Context) ([]domain.CarListResponse, error)
	Delete(c echo.Context, id uint64) error
	Patch(c echo.Context, req domain.CarPatchRequest) (domain.CarPatchResponse, error)
	Post(c echo.Context, req domain.CarPostRequest) (domain.CarPostResponse, error)
}

type CarController struct {
	cs CarService
}

func NewCarController(e *echo.Echo, cs CarService) {
	g := e.Group("/cars")
	cc := &CarController{
		cs: cs,
	}

	g.GET("", cc.List)
	g.DELETE("/:id", cc.Delete)
	g.PATCH("/:id", cc.Patch)
	g.POST("", cc.Post)
}

func (cc *CarController) Delete(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, echo.Map{
		"message": "not implemented",
	})
}

func (cc *CarController) Patch(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, echo.Map{
		"message": "not implemented",
	})
}

func (cc *CarController) Post(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, echo.Map{
		"message": "not implemented",
	})
}

func (cc *CarController) List(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, echo.Map{
		"message": "not implemented",
	})
}
