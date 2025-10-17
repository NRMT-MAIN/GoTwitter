-- +goose Up
-- +goose StatementBegin
CREATE TABLE tags (
	id SERIAL PRIMARY KEY, 
	name VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tags if exists;
-- +goose StatementEnd
