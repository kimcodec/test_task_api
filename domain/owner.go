package domain

type OwnerToStore struct {
	Name       string
	Surname    string
	Patronymic string
}

type OwnerDB struct {
	ID         uint64 `db:"id"`
	Name       string `db:"name"`
	Surname    string `db:"surname"`
	Patronymic string `db:"patronymic"`
}
