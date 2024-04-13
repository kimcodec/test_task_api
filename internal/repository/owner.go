package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kimcodec/test_api_task/domain"
	_ "github.com/lib/pq"
)

type OwnerRepository struct {
	db *sqlx.DB
}

func NewOwnerRepository(db *sqlx.DB) *OwnerRepository {
	return &OwnerRepository{
		db: db,
	}
}

func (or *OwnerRepository) Get(c context.Context, id uint64) (domain.OwnerDB, error) {
	return domain.OwnerDB{}, nil
}

func (or *OwnerRepository) Store(c context.Context, own domain.OwnerToStore) (domain.OwnerDB, error) {
	return domain.OwnerDB{}, nil
}
