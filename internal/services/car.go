package services

import (
	"context"
	"github.com/kimcodec/test_api_task/domain"
	openapi "github.com/kimcodec/test_api_task/outer_api"
	"log"
)

type CarRepository interface {
	Store(c context.Context, req domain.Car) (domain.CarDB, error)
	Delete(c context.Context, id uint64) error
	Patch(c context.Context, req domain.CarPatchRequest, id uint64) (domain.CarDB, error)
	List(c context.Context, params domain.CarFilterParams) ([]domain.CarWithOwnerDB, error)
}

type OwnerRepository interface {
	Get(c context.Context, id uint64) (domain.OwnerDB, error)
	Store(c context.Context, own domain.Owner) (domain.OwnerDB, error)
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

func (cs *CarService) Post(c context.Context, req domain.CarPostRequest) ([]domain.CarPostResponse, error) {
	var cars []domain.Car
	// Запрос к внешнему API
	for _, v := range req.RegNum {
		info := cs.outerApi.DefaultApi.InfoGet(c)
		info = info.RegNum(v)
		carResp, resp, err := info.Execute()
		if err != nil {
			log.Println("[ERROR] CarController.Post: Failed to get response from outer API: ", err.Error())
			return nil, err
		}

		car := domain.Car{
			Mark:   carResp.Mark,
			Model:  carResp.Model,
			RegNum: carResp.RegNum,
			Year:   carResp.Year,
			Owner: domain.Owner{
				Name:       carResp.Owner.Name,
				Surname:    carResp.Owner.Surname,
				Patronymic: carResp.Owner.Patronymic,
			},
		}

		cars = append(cars, car)
		resp.Body.Close()
	}

	var carsResp []domain.CarPostResponse
	for _, v := range cars {
		owDB, err := cs.or.Store(c, v.Owner)
		if err != nil {
			return nil, err
		}
		carDB, err := cs.cr.Store(c, v)
		if err != nil {
			return nil, err
		}

		car := domain.CarPostResponse{
			Model:  carDB.Model,
			Mark:   carDB.Mark,
			Year:   carDB.Year,
			RegNum: carDB.RegNum,
			Owner: domain.OwnerResponse{
				ID:         owDB.ID,
				Name:       owDB.Name,
				Surname:    owDB.Surname,
				Patronymic: owDB.Patronymic,
			},
		}
		carsResp = append(carsResp, car)
	}
	return carsResp, nil
}

func (cs *CarService) List(c context.Context, params domain.CarFilterParams) ([]domain.CarListResponse, error) {
	carsDB, err := cs.cr.List(c, params)
	if err != nil {
		return nil, err
	}

	var cars []domain.CarListResponse
	for _, v := range carsDB {
		car := domain.CarListResponse{
			ID:    v.ID,
			Mark:  v.Mark,
			Model: v.Model,
			Year:  v.Year,
			Owner: domain.OwnerResponse{
				ID:         v.Owner,
				Name:       v.Name,
				Surname:    v.Surname,
				Patronymic: v.Patronymic,
			},
		}

		cars = append(cars, car)
	}

	return cars, nil
}

func (cs *CarService) Delete(c context.Context, id uint64) error {
	if err := cs.cr.Delete(c, id); err != nil {
		return err
	}
	return nil
}

func (cs *CarService) Patch(c context.Context, req domain.CarPatchRequest, id uint64) (domain.CarPatchResponse, error) {
	carDB, err := cs.cr.Patch(c, req, id)
	if err != nil {
		return domain.CarPatchResponse{}, err
	}

	owDB, err := cs.or.Get(c, carDB.Owner)
	if err != nil {
		return domain.CarPatchResponse{}, err
	}

	car := domain.CarPatchResponse{
		ID:     carDB.ID,
		RegNum: carDB.RegNum,
		Mark:   carDB.Mark,
		Model:  carDB.Model,
		Year:   carDB.Year,
		Owner: domain.OwnerResponse{
			ID:         owDB.ID,
			Name:       owDB.Name,
			Surname:    owDB.Surname,
			Patronymic: owDB.Patronymic,
		},
	}

	return car, nil
}
