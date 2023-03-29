-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL,
	first_name  VARCHAR(20) NOT NULL,
	last_name  VARCHAR(20) NOT NULL,
	username  VARCHAR(20) NOT NULL,
	email  VARCHAR(50) NOT NULL,
	password TEXT NOT NULL,
	is_admin BOOLEAN DEFAULT FALSE,
	is_active BOOLEAN DEFAULT TRUE,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id),
	UNIQUE(username),
	UNIQUE(email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users ;
-- +goose StatementEnd
