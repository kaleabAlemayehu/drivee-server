-- +goose Up
-- +goose StatementBegin
CREATE TYPE PAYMENT_STATUS AS ENUM('pending', 'paid', 'failed', 'refunded');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS PAYMENT_STATUS CASCADE;
-- +goose StatementEnd
