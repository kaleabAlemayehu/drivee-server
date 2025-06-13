-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  first_name TEXT NOT NULL,
  middle_name TEXT,
  last_name TEXT ,
  email TEXT NOT NULL UNIQUE,
  profile_picture TEXT NOT NULL,
  password TEXT,
  driver_license  TEXT,
  phone_number TEXT,
  account_number TEXT,
  bank_name TEXT,
  is_owner BOOLEAN DEFAULT false,
  is_renter BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd
