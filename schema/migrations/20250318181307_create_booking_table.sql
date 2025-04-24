-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bookings (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  car_id UUID , 
	-- TODO:( fix this on prod) REFERENCES cars(id) NOT NULL,
  renter_id UUID NOT NULL,
	-- TODO:( fix this on prod)  REFERENCES users(id) NOT NULL,
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  total_price NUMERIC(10, 2) NOT NULL,
  status BOOKING_STATUS NOT NULL DEFAULT 'pending',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bookings CASCADE;
-- +goose StatementEnd
