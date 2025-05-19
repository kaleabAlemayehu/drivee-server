-- +goose Up
-- +goose StatementBegin
-- INSERT INTO users ( first_name, middle_name, last_name, email, password, driver_license, phone_number, account_number, bank_name) VALUES ('Elliot','Mr.Robot', 'Alderson', 'Elliot1@gmail.com', 'password', 'his driving license', '+10232023432', '397519708', 'BOA');
-- INSERT INTO users ( first_name, middle_name, last_name, email, password, driver_license, phone_number, account_number, bank_name) VALUES ('Elliot','Mr.Robot', 'Alderson', 'Elliot2@gmail.com', 'password', 'his driving license', '+10232023432', '397519708', 'BOA');
-- INSERT INTO users ( first_name, middle_name, last_name, email, password, driver_license, phone_number, account_number, bank_name) VALUES ('Elliot','Mr.Robot', 'Alderson', 'Elliot3@gmail.com', 'password', 'his driving license', '+10232023432', '397519708', 'BOA');
-- INSERT INTO users ( first_name, middle_name, last_name, email, password, driver_license, phone_number, account_number, bank_name) VALUES ('Elliot','Mr.Robot', 'Alderson', 'Elliot4@gmail.com', 'password', 'his driving license', '+10232023432', '397519708', 'BOA');
-- INSERT INTO users ( first_name, middle_name, last_name, email, password, driver_license, phone_number, account_number, bank_name) VALUES ('Elliot','Mr.Robot', 'Alderson', 'Elliot5@gmail.com', 'password', 'his driving license', '+10232023432', '397519708', 'BOA');
-- INSERT INTO users ( first_name, middle_name, last_name, email, password, driver_license, phone_number, account_number, bank_name) VALUES ('Elliot','Mr.Robot', 'Alderson', 'Elliot6@gmail.com', 'password', 'his driving license', '+10232023432', '397519708', 'BOA');
-- INSERT INTO users ( first_name, middle_name, last_name, email, password, driver_license, phone_number, account_number, bank_name) VALUES ('Elliot','Mr.Robot', 'Alderson', 'Elliot7@gmail.com', 'password', 'his driving license', '+10232023432', '397519708', 'BOA');
-- INSERT INTO users ( first_name, middle_name, last_name, email, password, driver_license, phone_number, account_number, bank_name) VALUES ('Elliot','Mr.Robot', 'Alderson', 'Elliot8@gmail.com', 'password', 'his driving license', '+10232023432', '397519708', 'BOA');
-- INSERT INTO users ( first_name, middle_name, last_name, email, password, driver_license, phone_number, account_number, bank_name) VALUES ('Elliot','Mr.Robot', 'Alderson', 'Elliot9@gmail.com', 'password', 'his driving license', '+10232023432', '397519708', 'BOA');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE users CASCADE;
-- +goose StatementEnd
