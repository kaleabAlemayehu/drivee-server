-- +goose Up
-- +goose StatementBegin
INSERT INTO owner ( first_name, last_name, email, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', '+10232023432', '397519708', 'BOA');
INSERT INTO owner ( first_name, last_name, email, phone_number, account_number, bank_name) VALUES ('Elliot', 'Alderson', 'Elliot@gmail.com', '+10232023432', '397519708', 'BOA');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM owner WHERE first_name='kaleab';
-- +goose StatementEnd
