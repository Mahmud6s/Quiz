-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quiz_questions (
	id BIGSERIAL,
	quiz_id  INT NOT NULL,
	question_id  INT NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(quiz_id) REFERENCES quiz(id),
	FOREIGN KEY(question_id) REFERENCES questions(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quiz_questions;
-- +goose StatementEnd
