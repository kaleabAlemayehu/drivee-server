-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cars (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  owner_id UUID,
-- TODO: (fix this on production) REFERENCES users(id) NOT NULL,
  thumbnail_picture TEXT NOT NULL,
  description TEXT NOT NULL,
  make TEXT NOT NULL,
  model TEXT NOT NULL,
  year TEXT NOT NULL,
  license_plate TEXT UNIQUE NOT NULL,
  vin_number TEXT UNIQUE NOT NULL,
  transmission TRANSMISSION  NOT NULL,
  fuel_type FUEL_TYPE NOT NULL,
  mileage INT NOT NULL,
  location GEOMETRY(POINT, 4326) NOT NULL,
  price_per_hour NUMERIC(10,2) NOT NULL,
  status STATUS_TYPE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cars CASCADE;
-- +goose StatementEnd
