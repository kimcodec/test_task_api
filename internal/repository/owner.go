package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/kimcodec/test_api_task/domain"

	"context"
	"errors"
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
		return domain.OwnerDB{}, err
	}
	defer conn.Close()

	var own []domain.OwnerDB
	if err := conn.SelectContext(c, &own, "SELECT * FROM Owners WHERE id = $1", id); err != nil {
		log.Println("[ERROR] OwnerRepository.Get: failed to execute query: ", err.Error())
		return domain.OwnerDB{}, err
	}

	if len(own) == 0 {
		return domain.OwnerDB{}, errors.New("owner does not exists")
	}

	return own[0], nil
}

func (or *OwnerRepository) Store(c context.Context, own domain.Owner) (domain.OwnerDB, error) {
	conn, err := or.db.Connx(c)
	if err != nil {
		log.Println("[ERROR] OwnerRepository.Store: failed to create connection to db: ", err.Error())
		return domain.OwnerDB{}, err
	}
	defer conn.Close()

	row := conn.QueryRowxContext(c,
		"INSERT INTO Owners(name, surname, patronymic) VALUES($1, $2, $3) RETURNING *",
		own.Name, own.Surname, own.Patronymic)
	if row.Err() != nil {
		log.Println("[ERROR] OwnerRepository.Store: failed to execute query: ", row.Err().Error())
		return domain.OwnerDB{}, row.Err()
	}

	var owner domain.OwnerDB
	if err := row.StructScan(&owner); err != nil {
		log.Println("[ERROR] OwnerRepository.Store: failed to scan struct: ", err.Error())
		return domain.OwnerDB{}, err
	}

	return owner, nil
}
