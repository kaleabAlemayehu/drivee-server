-- +goose Up
-- +goose StatementBegin
CREATE TABLE car_photos(
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  car_id UUID,
-- TODO: change this on production REFERENCES cars(id) NOT NULL,
  photo_url TEXT NOT NULL,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS car_photos CASCADE;
-- +goose StatementEnd
