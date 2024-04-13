package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kimcodec/test_api_task/domain"
	_ "github.com/lib/pq"
	"log"
)

type CarRepository struct {
	db *sqlx.DB
}

func NewCarRepository(db *sqlx.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (cr *CarRepository) Store(c context.Context, req domain.Car, ownerID uint64) (domain.CarDB, error) {
	conn, err := cr.db.Connx(c)
	if err != nil {
		log.Println("[ERROR] CarRepository.Store: failed to create connection to db: ", err.Error())
		return domain.CarDB{}, nil
	}
	defer conn.Close()

	var car domain.CarDB
	row, err := conn.QueryxContext(c,
		"INSERT INTO Cars(owner_id, reg_num, mark, model, year) "+
			"VALUES ($1, $2, $3, $4, $5) RETURNING *", ownerID, req.RegNum, req.Mark, req.Model, req.Year)
	if err != nil {
		log.Println("[ERROR] CarRepository.Store: failed to execute row: ", err.Error())
		return domain.CarDB{}, err
	}
	defer row.Close()

	if err := row.StructScan(&car); err != nil {
		log.Println("[ERROR] CarRepository.Store: failed to scan struct: ", err.Error())
		return domain.CarDB{}, err
	}

	return car, nil
}

func (cr *CarRepository) Delete(c context.Context, id uint64) error {
	conn, err := cr.db.Connx(c)
	if err != nil {
		log.Println("[ERROR] CarRepository.Delete: failed to create connection to db: ", err.Error())
		return nil
	}
	defer conn.Close()

	if _, err := conn.ExecContext(c, "DELETE FROM Cars WHERE id = $1", id); err != nil {
		log.Println("[ERROR] CarRepository.Delete: failed to execute query: ", err.Error())
		return err
	}

	return nil
}

func (cr *CarRepository) Patch(c context.Context, req domain.CarPatchRequest, id uint64) (domain.CarDB, error) {
	conn, err := cr.db.Connx(c)
	if err != nil {
		log.Println("[ERROR] CarRepository.Patch: failed to create connection to db: ", err.Error())
		return domain.CarDB{}, nil
	}
	defer conn.Close()

	row, err := conn.QueryxContext(c,
		"UPDATE Cars SET RegNum = $1, Mark = $2, Model = $3, Year = $4 WHERE id = $5 RETURNING *",
		req.RegNum, req.Mark, req.Model, req.Year, id)
	if err != nil {
		log.Println("[ERROR] CarRepository.Patch: failed execute query: ", err.Error())
		return domain.CarDB{}, err
	}
	defer row.Close()

	var car domain.CarDB
	if err := row.StructScan(&car); err != nil {
		log.Println("[ERROR] CarRepository.Patch: failed to scan struct: ", err.Error())
		return domain.CarDB{}, err
	}
	return car, nil
}

func (cr *CarRepository) List(c context.Context, params domain.CarFilterParams) ([]domain.CarWithOwnerDB, error) {
	conn, err := cr.db.Connx(c)
	if err != nil {
		log.Println("[ERROR] CarRepository.List: failed to create connection to db: ", err.Error())
		return nil, nil
	}
	defer conn.Close()

	var carsWithOwners []domain.CarWithOwnerDB
	if err := conn.SelectContext(c, &carsWithOwners,
		"SELECT * FROM Cars JOIN Owners ON Cars.owner_id = Owners.id "+
			"WHERE Cars.id >= $1 AND year >= $2 AND mark LIKE $3 AND model LIKE $4 AND reg_num LIKE $5 LIMIT $6",
		params.Offset, params.Year, params.Mark, params.Model, params.RegNum, params.Limit); err != nil {
		log.Println("[ERROR] CarRepository.List: failed to execute query: ", err.Error())
		return nil, err
	}

	return carsWithOwners, nil
}
