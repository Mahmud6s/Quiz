-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_quiz (
	id BIGSERIAL,
	users_id  INT NOT NULL,
	quizquestion_id  INT NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(users_id) REFERENCES users(id),
	FOREIGN KEY(quizquestion_id) REFERENCES quiz_questions(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_quiz;
-- +goose StatementEnd
