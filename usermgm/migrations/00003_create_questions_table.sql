-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS questions (
	id BIGSERIAL,
	category_id  INT NOT NULL,
	title  TEXT NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(category_id) REFERENCES categorys(id)
	
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS questions;
-- +goose StatementEnd
