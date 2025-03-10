-- +goose Up
-- +goose StatementBegin
CREATE TYPE TRANSMISSION AS ENUM('manual', 'automatic');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS TRANSMISSION CASCADE;
-- +goose StatementEnd
