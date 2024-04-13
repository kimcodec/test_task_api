package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kimcodec/test_api_task/domain"
	_ "github.com/lib/pq"
	"log"
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
	conn, err := or.db.Connx(c)
	if err != nil {
		log.Println("[ERROR] OwnerRepository.Get: failed to create connection to db: ", err.Error())
		return domain.OwnerDB{}, nil
	}
	defer conn.Close()

	var own domain.OwnerDB
	if err := conn.SelectContext(c, &own, "SELECT * FROM Owners WHERE id = $1", id); err != nil {
		log.Println("[ERROR] OwnerRepository.Get: failed to execute query: ", err.Error())
		return domain.OwnerDB{}, err
	}

	return own, nil
}

func (or *OwnerRepository) Store(c context.Context, own domain.Owner) (domain.OwnerDB, error) {
	conn, err := or.db.Connx(c)
	if err != nil {
		log.Println("[ERROR] OwnerRepository.Store: failed to create connection to db: ", err.Error())
		return domain.OwnerDB{}, nil
	}
	defer conn.Close()

	row, err := conn.QueryxContext(c,
		"INSERT INTO Owners(name, surname, patronymic) VALUES($1, $2, $3) RETURNING *",
		own.Name, own.Surname, own.Patronymic)
	if err != nil {
		log.Println("[ERROR] OwnerRepository.Store: failed to execute query: ", err.Error())
		return domain.OwnerDB{}, err
	}

	var owner domain.OwnerDB
	if err := row.StructScan(&owner); err != nil {
		log.Println("[ERROR] OwnerRepository.Store: failed to scan struct: ", err.Error())
		return domain.OwnerDB{}, err
	}

	return owner, nil
}
