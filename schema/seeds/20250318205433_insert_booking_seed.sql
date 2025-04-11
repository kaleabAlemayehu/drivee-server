-- +goose Up
-- +goose StatementBegin
INSERT INTO bookings (car_id, renter_id, start_time, end_time, total_price, status) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','0384ff04-079c-4d03-a205-370f3a9627ac', now(), now(), 100,'pending');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE cars CASCADE;
-- +goose StatementEnd
