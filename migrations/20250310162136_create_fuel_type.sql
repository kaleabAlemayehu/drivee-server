-- +goose Up
-- +goose StatementBegin
CREATE TYPE FUEL_TYPE AS ENUM('gas', 'diesel', 'electric', 'hybrid');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS FUEL_TYPE CASCADE;
-- +goose StatementEnd
