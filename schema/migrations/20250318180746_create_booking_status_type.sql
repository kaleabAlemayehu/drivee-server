-- +goose Up
-- +goose StatementBegin
CREATE TYPE BOOKING_STATUS AS ENUM('pending', 'confirmed', 'canceled', 'completed');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS BOOKING_STATUS CASCADE;
-- +goose StatementEnd
