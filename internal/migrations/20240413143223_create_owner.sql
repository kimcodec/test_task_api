-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Owners(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname varchar(255) NOT NULL,
    patronymic varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Owners;
-- +goose StatementEnd
