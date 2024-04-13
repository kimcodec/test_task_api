-- +goose Up
-- +goose StatementBegin
CREATE TABLE Cars(
    id SERIAL PRIMARY KEY,
    owner_id INT NOT NULL,
    reg_num VARCHAR(256) NOT NULL,
    mark VARCHAR(256) NOT NULL,
    model VARCHAR(256) NOT NULL,
    year INT NOT NULL,
    CONSTRAINT fk_cars_owner_id
    FOREIGN KEY (owner_id)
        REFERENCES Owners(id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS CARS;
-- +goose StatementEnd
