-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  booking_id UUID,
-- TODO: (fix this on production) REFERENCES bookings(id) NOT NULL
  renter_id UUID,
-- TODO: (fix this on production) REFERENCES users(id) NOT NULL
  owner_id UUID,
-- TODO: (fix this on production) REFERENCES users(id) NOT NULL
  amount NUMERIC(10,2) NOT NULL,
  payment_status PAYMENT_STATUS NOT NULL,
  payment_method PAYMENT_METHOD NOT NULL,
  transaction_id VARCHAR(255) UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments CASCADE;
-- +goose StatementEnd
