-- +goose Up
-- +goose StatementBegin
CREATE TYPE STATUS_TYPE AS ENUM('avaliable', 'rented', 'inactive');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS STATUS_TYPE CASCADE;
-- +goose StatementEnd
