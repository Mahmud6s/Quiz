-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quiz (
	id BIGSERIAL,
	category_id  INT NOT NULL,
	quiz_title  TEXT NOT NULL,
	quiz_time  INT NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(category_id) REFERENCES categorys(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quiz;
-- +goose StatementEnd
