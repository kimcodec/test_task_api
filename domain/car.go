package domain

type CarFilterParams struct {
	Offset uint64
	Limit  uint64
	Year   uint64
	Mark   string
	Model  string
	RegNum string
}

type CarPostRequest struct {
	RegNums []string `json:"reg_nums"`
}

type CarPostResponse struct {
	ID     uint64        `json:"id"`
	RegNum string        `json:"reg_num"`
	Mark   string        `json:"mark"`
	Model  string        `json:"model"`
	Year   *int32        `json:"year,omitempty"`
	Owner  OwnerResponse `json:"owner"`
}

type CarListResponse struct {
	ID     uint64        `json:"id"`
	RegNum string        `json:"reg_num"`
	Mark   string        `json:"mark"`
	Model  string        `json:"model"`
	Year   *int32        `json:"year,omitempty"`
	Owner  OwnerResponse `json:"owner"`
}

type CarPatchRequest struct {
	RegNum *string `json:"reg_num"`
	Mark   *string `json:"mark"`
	Model  *string `json:"model"`
	Year   *string `json:"year"`
}

type CarPatchResponse struct {
	ID     uint64        `json:"id"`
	RegNum string        `json:"reg_num"`
	Mark   string        `json:"mark"`
	Model  string        `json:"model"`
	Year   *int32        `json:"year,omitempty"`
	Owner  OwnerResponse `json:"owner"`
}

type Car struct {
	RegNum string
	Mark   string
	Model  string
	Year   *int32
	Owner  Owner
}

type CarDB struct {
	ID     uint64 `db:"id"`
	Owner  uint64 `db:"owner_id"`
	Year   *int32 `db:"year"`
	RegNum string `db:"reg_num"`
	Mark   string `db:"mark"`
	Model  string `db:"model"`
}

type CarWithOwnerDB struct {
	ID         uint64  `db:"id"`
	Owner      uint64  `db:"owner_id"`
	Name       string  `db:"name"`
	Surname    string  `db:"surname"`
	Patronymic *string `db:"patronymic"`
	Year       *int32  `db:"year"`
	RegNum     string  `db:"reg_num"`
	Mark       string  `db:"mark"`
	Model      string  `db:"model"`
}
