-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS answer (
	id BIGSERIAL,
	userquiz_id  INT NOT NULL,
	question_id  INT NOT NULL,
	option_id  INT NOT NULL,
	is_correct BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(userquiz_id) REFERENCES user_quiz(id),
	FOREIGN KEY(question_id) REFERENCES questions(id),
	FOREIGN KEY(option_id) REFERENCES options(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS answer;
-- +goose StatementEnd
