-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reviews(
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  reviewer_id UUID,
-- TODO: (fix this on production) REFERENCES users(id) NOT NULL
  target_id UUID,
-- TODO: (fix this on production) REFERENCES users(id) NOT NULL
  booking_id UUID,
-- TODO: (fix this on production) REFERENCES bookings(id) NOT NULL
  rating INT CHECK(rating <= 5 AND rating >= 1) NOT NULL,
  comment TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reviews CASCADE;
-- +goose StatementEnd
