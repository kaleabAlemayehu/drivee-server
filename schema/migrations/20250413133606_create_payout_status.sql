-- +goose Up
-- +goose StatementBegin
CREATE TYPE PAYOUT_STATUS AS ENUM('pending', 'completed', 'failed');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS PAYOUT_STATUS CASCADE;
-- +goose StatementEnd
