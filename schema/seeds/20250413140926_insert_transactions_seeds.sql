-- +goose Up
-- +goose StatementBegin
INSERT INTO transactions(renter_id, owner_id, booking_id, amount, fee, payout_status) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe', '5d3bd16e-b185-44cc-ae48-76b86532cffe','5d3bd16e-b185-44cc-ae48-76b86532cffe', 1000, 50, 'pending' );
INSERT INTO transactions(renter_id, owner_id, booking_id, amount, fee, payout_status) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe', '5d3bd16e-b185-44cc-ae48-76b86532cffe','5d3bd16e-b185-44cc-ae48-76b86532cffe', 1000, 50, 'pending' );
INSERT INTO transactions(renter_id, owner_id, booking_id, amount, fee, payout_status) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe', '5d3bd16e-b185-44cc-ae48-76b86532cffe','5d3bd16e-b185-44cc-ae48-76b86532cffe', 1000, 50, 'pending' );
INSERT INTO transactions(renter_id, owner_id, booking_id, amount, fee, payout_status) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe', '5d3bd16e-b185-44cc-ae48-76b86532cffe','5d3bd16e-b185-44cc-ae48-76b86532cffe', 1000, 50, 'pending' );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE transactions CASCADE;
-- +goose StatementEnd
