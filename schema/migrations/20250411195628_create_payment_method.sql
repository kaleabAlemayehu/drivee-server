-- +goose Up
-- +goose StatementBegin
CREATE TYPE PAYMENT_METHOD AS ENUM('credit_card', 'paypal', 'apple_pay','telebirr','google_pay');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS PAYMENT_METHOD CASCADE;
-- +goose StatementEnd
