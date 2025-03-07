-- +goose Up
-- +goose StatementBegin
INSERT INTO owner ( first_name, last_name, email, password, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', 'password', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, password, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', 'password', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, password, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', 'password', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, password, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', 'password', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, password, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', 'password', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, password, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', 'password', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, password, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', 'password', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, password, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', 'password', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, password, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', 'password', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, password, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', 'password', '+10232023432', '397519708', 'BOA');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM owner WHERE bank_name='BOA';
-- +goose StatementEnd
