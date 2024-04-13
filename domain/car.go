package domain

import openapi "github.com/kimcodec/test_api_task/outer_api"

type CarFilterParams struct {
}

type CarPostRequest struct {
	RegNum []string `json:"reg_num"`
}

type CarPostResponse struct {
	ID     uint64         `json:"id"`
	RegNum string         `json:"reg_num"`
	Mark   string         `json:"mark"`
	Model  string         `json:"model"`
	Owner  openapi.People `json:"owner"`
}

type CarListResponse struct {
	ID     uint64         `json:"id"`
	RegNum string         `json:"reg_num"`
	Mark   string         `json:"mark"`
	Model  string         `json:"model"`
	Owner  openapi.People `json:"owner"`
}

type CarPatchRequest struct {
	ID     uint64 `json:"id"`
	RegNum string `json:"reg_num"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
}

type CarPatchResponse struct {
	ID     uint64         `json:"id"`
	RegNum string         `json:"reg_num"`
	Mark   string         `json:"mark"`
	Model  string         `json:"model"`
	Owner  openapi.People `json:"owner"`
}

type CarDB struct {
	ID     uint64 `db:"id"`
	RegNum string `db:"reg_num"`
	Mark   string `db:"mark"`
	Model  string `db:"model"`
	Owner  uint64 `db:"owner_id"`
}
