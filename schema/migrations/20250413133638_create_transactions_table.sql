-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transactions (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  renter_id UUID,
-- TODO: (fix this on production) REFERENCES users(id) NOT NULL
  owner_id UUID,
-- TODO: (fix this on production) REFERENCES users(id) NOT NULL
  booking_id UUID,
-- TODO: (fix this on production) REFERENCES bookings(id) NOT NULL
  amount NUMERIC(10, 2) NOT NULL,
  fee NUMERIC(10, 2) NOT NULL,
  payout_status PAYOUT_STATUS NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions CASCADE;
-- +goose StatementEnd
