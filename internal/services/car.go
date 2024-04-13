package services

import (
	"context"
	"github.com/kimcodec/test_api_task/domain"
	openapi "github.com/kimcodec/test_api_task/outer_api"
)

type CarRepository interface {
	Store(c context.Context, req domain.CarPostRequest) (domain.CarDB, error)
	Delete(c context.Context, id uint64) error
	Patch(c context.Context, req domain.CarPatchRequest) (domain.CarDB, error)
	List(c context.Context, params domain.CarFilterParams) ([]domain.CarDB, error)
}

type OwnerRepository interface {
	Get(c context.Context, id uint64) (domain.OwnerDB, error)
	Store(c context.Context, own domain.OwnerToStore) (domain.OwnerDB, error)
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

func (cr *CarService) List(c context.Context, params domain.CarFilterParams) ([]domain.CarListResponse, error) {
	return nil, nil
}

func (cr *CarService) Delete(c context.Context, id uint64) error {
	return nil
}

func (cr *CarService) Patch(c context.Context, req domain.CarPatchRequest, id uint64) (domain.CarPatchResponse, error) {
	return domain.CarPatchResponse{}, nil
}

func (cr *CarService) Post(c context.Context, req domain.CarPostRequest) (domain.CarPostResponse, error) {
	return domain.CarPostResponse{}, nil
}
