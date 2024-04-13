package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kimcodec/test_api_task/domain"
	_ "github.com/lib/pq"
)

type CarRepository struct {
	db *sqlx.DB
}

func NewCarRepository(db *sqlx.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (cr *CarRepository) Store(c context.Context, req domain.Car) (domain.CarDB, error) {
	return domain.CarDB{}, nil
}

func (cr *CarRepository) Delete(c context.Context, id uint64) error {
	return nil
}

func (cr *CarRepository) Patch(c context.Context, req domain.CarPatchRequest, id uint64) (domain.CarDB, error) {
	return domain.CarDB{}, nil
}

func (cr *CarRepository) List(c context.Context, params domain.CarFilterParams) ([]domain.CarWithOwnerDB, error) {
	return nil, nil
}
