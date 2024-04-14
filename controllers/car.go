package controllers

import (
	"github.com/labstack/echo/v4"

	"github.com/kimcodec/test_api_task/domain"

	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type CarService interface {
	List(c context.Context, params domain.CarFilterParams) ([]domain.CarListResponse, error)
	Delete(c context.Context, id uint64) error
	Patch(c context.Context, req domain.CarPatchRequest, id uint64) (domain.CarPatchResponse, error)
	Post(c context.Context, req domain.CarPostRequest) ([]domain.CarPostResponse, error)
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

func (cc *CarController) Post(c echo.Context) error {
	var carReq domain.CarPostRequest
	if err := c.Bind(&carReq); err != nil {
		log.Println("[ERROR] CarController.Post: Failed to parse JSON : ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": fmt.Sprintf("Failed to parse JSON: %s", err.Error()),
		})
	}

	ctx := c.Request().Context()

	car, err := cc.cs.Post(ctx, carReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("Failed to add car: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, car)
}

func (cc *CarController) Delete(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		log.Println("[ERROR] CarController.Delete: Failed to get id from url path")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to get id",
		})
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("[ERROR] CarController.Delete: Failed to parse id: ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": fmt.Sprintf("failed to parse id: %s", err.Error()),
		})
	}

	ctx := c.Request().Context()

	if err := cc.cs.Delete(ctx, uint64(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("failed to delete car %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "car was deleted",
	})
}

func (cc *CarController) Patch(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		log.Println("[ERROR] CarController.Patch: Failed to get id from url path")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to get id",
		})
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("[ERROR] CarController.Delete: Failed to parse id: ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": fmt.Sprintf("failed to parse id: %s", err.Error()),
		})
	}

	var carReq domain.CarPatchRequest
	if err := c.Bind(&carReq); err != nil {
		log.Println("[ERROR] CarController.Patch: Failed to parse JSON : ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": fmt.Sprintf("Failed to parse JSON: %s", err.Error()),
		})
	}

	ctx := c.Request().Context()

	car, err := cc.cs.Patch(ctx, carReq, uint64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("failed to update car: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, car)
}

func (cc *CarController) List(c echo.Context) error {
	params, err := getQueryParams(c)
	if err != nil {
		log.Println("[ERROR] CarController.Delete: Failed to get params: ", err.Error())
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": fmt.Sprintf("failed to get params: %s", err.Error()),
		})
	}

	ctx := c.Request().Context()
	cars, err := cc.cs.List(ctx, params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("failed to get cars: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, cars)
}

func getQueryParams(c echo.Context) (domain.CarFilterParams, error) {
	var (
		limit  uint64
		offset uint64
		year   uint64
		mark   string
		model  string
		regNum string
	)

	limitStr := c.QueryParam("limit")
	if limitStr != "" {
		limitTemp, err := strconv.Atoi(limitStr)
		if err != nil {
			return domain.CarFilterParams{}, err
		}
		limit = uint64(limitTemp)
	} else {
		limit = 10
	}

	offsetStr := c.QueryParam("offset")
	if offsetStr != "" {
		offsetTemp, err := strconv.Atoi(offsetStr)
		if err != nil {
			return domain.CarFilterParams{}, err
		}
		offset = uint64(offsetTemp)
	} else {
		offset = 0
	}

	yearStr := c.QueryParam("year")
	if yearStr != "" {
		yearTemp, err := strconv.Atoi(yearStr)
		if err != nil {
			return domain.CarFilterParams{}, err
		}
		year = uint64(yearTemp)
	} else {
		year = 0
	}

	mark = c.QueryParam("mark")
	if mark == "" {
		mark = "%"
	} else {
		mark = "%" + mark + "%"
	}

	model = c.QueryParam("model")
	if model == "" {
		model = "%"
	} else {
		model = "%" + model + "%"
	}

	regNum = c.QueryParam("reg_num")
	if regNum == "" {
		regNum = "%"
	} else {
		regNum = "%" + regNum + "%"
	}

	return domain.CarFilterParams{
		Offset: offset,
		Limit:  limit,
		Year:   year,
		Model:  model,
		Mark:   mark,
		RegNum: regNum,
	}, nil
}
