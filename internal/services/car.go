package services

import (
	"github.com/kimcodec/test_api_task/domain"
	openapi "github.com/kimcodec/test_api_task/outer_api"
	"github.com/labstack/echo/v4"
)

type CarRepository interface {
	Store(c echo.Context, req domain.CarPostRequest) (domain.CarDB, error)
	Delete(c echo.Context, id uint64) error
	Patch(c echo.Context, req domain.CarPatchRequest) (domain.CarDB, error)
	List(c echo.Context, params domain.CarFilterParams) ([]domain.CarDB, error)
}

type OwnerRepository interface {
	Get(c echo.Context, id uint64) (domain.OwnerDB, error)
	Store(c echo.Context, own domain.OwnerToStore) (domain.OwnerDB, error)
}

type CarService struct {
	cr       CarRepository
	or       OwnerRepository
	outerApi *openapi.APIClient
}

func NewCarService(cr CarRepository, or OwnerRepository, outerApi *openapi.APIClient) *CarService {
	return &CarService{
		cr:       cr,
		or:       or,
		outerApi: outerApi,
	}
}

func (cr *CarService) List(c echo.Context) ([]domain.CarListResponse, error) {
	return nil, nil
}

func (cr *CarService) Delete(c echo.Context, id uint64) error {
	return nil
}

func (cr *CarService) Patch(c echo.Context, req domain.CarPatchRequest) (domain.CarPatchResponse, error) {
	return domain.CarPatchResponse{}, nil
}

func (cr *CarService) Post(c echo.Context, req domain.CarPostRequest) (domain.CarPostResponse, error) {
	return domain.CarPostResponse{}, nil
}
