-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS owner (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  first_name VARCHAR(24) NOT NULL,
  middle_name VARCHAR(24),
  last_name VARCHAR(24) NOT NULL,
  email VARCHAR(50) NOT NULL,
  password TEXT NOT NULL,
  phone_number VARCHAR(15) NOT NULL,
  account_number VARCHAR(15) NOT NULL,
  bank_name VARCHAR(15) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS owner CASCADE;
-- +goose StatementEnd
