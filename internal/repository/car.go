package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/kimcodec/test_api_task/domain"

	"context"
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
		return domain.CarDB{}, err
	}
	defer conn.Close()

	var car domain.CarDB
	row := conn.QueryRowxContext(c,
		"INSERT INTO Cars(owner_id, reg_num, mark, model, year) "+
			"VALUES ($1, $2, $3, $4, $5) RETURNING *", ownerID, req.RegNum, req.Mark, req.Model, req.Year)
	if row.Err() != nil {
		log.Println("[ERROR] CarRepository.Store: failed to execute row: ", row.Err().Error())
		return domain.CarDB{}, row.Err()
	}

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
		return err
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
		return domain.CarDB{}, err
	}
	defer conn.Close()

	row := conn.QueryRowxContext(c,
		"UPDATE Cars SET reg_num = COALESCE($1, reg_num), mark = COALESCE($2, mark), "+
			"model = COALESCE($3, model), year = COALESCE($4, year) WHERE id = $5 RETURNING *",
		req.RegNum, req.Mark, req.Model, req.Year, id)
	if row.Err() != nil {
		log.Println("[ERROR] CarRepository.Patch: failed execute query: ", row.Err().Error())
		return domain.CarDB{}, row.Err()
	}

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
		return nil, err
	}
	defer conn.Close()

	var carsWithOwners []domain.CarWithOwnerDB
	if err := conn.SelectContext(c, &carsWithOwners,
		"SELECT cars.id, owner_id, name, surname, patronymic, year, reg_num, mark, model "+
			"FROM Cars JOIN Owners ON Cars.owner_id = Owners.id "+
			"WHERE year >= $1 AND mark LIKE $2 AND model LIKE $3 AND reg_num LIKE $4 ORDER BY Cars.id LIMIT $5 OFFSET $6",
		params.Year, params.Mark, params.Model, params.RegNum, params.Limit, params.Offset); err != nil {
		log.Println("[ERROR] CarRepository.List: failed to execute query: ", err.Error())
		return nil, err
	}
	log.Println(carsWithOwners)

	return carsWithOwners, nil
}
