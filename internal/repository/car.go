package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kimcodec/test_api_task/domain"
	"github.com/labstack/echo/v4"
)

type CarRepository struct {
	db *sqlx.DB
}

func NewCarRepository(db *sqlx.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (cr *CarRepository) Store(c echo.Context, req domain.CarPostRequest) (domain.CarDB, error) {
	return domain.CarDB{}, nil
}

func (cr *CarRepository) Delete(c echo.Context, id uint64) error {
	return nil
}

func (cr *CarRepository) Patch(c echo.Context, req domain.CarPatchRequest) (domain.CarDB, error) {
	return domain.CarDB{}, nil
}

func (cr *CarRepository) List(c echo.Context, params domain.CarFilterParams) ([]domain.CarDB, error) {
	return nil, nil
}
