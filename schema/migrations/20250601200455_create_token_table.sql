-- +goose Up
-- +goose StatementBegin
CREATE TABLE token (
	token TEXT NOT NULL UNIQUE,
	expires_at INTEGER NOT NULL,
	user_id UUID NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS token CASCADE;
-- +goose StatementEnd
