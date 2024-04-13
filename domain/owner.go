package domain

type OwnerResponse struct {
	ID         uint64  `json:"id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
}

type Owner struct {
	Name       string
	Surname    string
	Patronymic *string
}

type OwnerDB struct {
	ID         uint64  `db:"id"`
	Name       string  `db:"name"`
	Surname    string  `db:"surname"`
	Patronymic *string `db:"patronymic"`
}
