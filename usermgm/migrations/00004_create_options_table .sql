-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS options (
	id BIGSERIAL,
	question_id  INT NOT NULL,
	option_name  TEXT NOT NULL,
	is_correct BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(question_id) REFERENCES questions(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS options;
-- +goose StatementEnd
