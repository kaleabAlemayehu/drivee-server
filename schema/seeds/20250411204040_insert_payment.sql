-- +goose Up
-- +goose StatementBegin
INSERT INTO payments(booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','0384ff04-079c-4d03-a205-370f3a9627ac', '0384ff04-079c-4d03-a205-370f3a9627ac', 2000, 'pending', 'paypal', 'testpayment0');
INSERT INTO payments(booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','0384ff04-079c-4d03-a205-370f3a9627ac', '0384ff04-079c-4d03-a205-370f3a9627ac', 2000, 'pending', 'paypal', 'testpayment1');
INSERT INTO payments(booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','0384ff04-079c-4d03-a205-370f3a9627ac', '0384ff04-079c-4d03-a205-370f3a9627ac', 2000, 'pending', 'paypal', 'testpayment2');
INSERT INTO payments(booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','0384ff04-079c-4d03-a205-370f3a9627ac', '0384ff04-079c-4d03-a205-370f3a9627ac', 2000, 'pending', 'paypal', 'testpayment3');
INSERT INTO payments(booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','0384ff04-079c-4d03-a205-370f3a9627ac', '0384ff04-079c-4d03-a205-370f3a9627ac', 2000, 'pending', 'paypal', 'testpayment4');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE payments CASCADE;
-- +goose StatementEnd
