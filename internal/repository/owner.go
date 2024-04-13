package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kimcodec/test_api_task/domain"
	"github.com/labstack/echo/v4"
)

type OwnerRepository struct {
	db *sqlx.DB
}

func NewOwnerRepository(db *sqlx.DB) *OwnerRepository {
	return &OwnerRepository{
		db: db,
	}
}

func (or *OwnerRepository) Get(c echo.Context, id uint64) (domain.OwnerDB, error) {
	return domain.OwnerDB{}, nil
}

func (or *OwnerRepository) Store(c echo.Context, own domain.OwnerToStore) (domain.OwnerDB, error) {
	return domain.OwnerDB{}, nil
}
