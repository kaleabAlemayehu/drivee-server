-- +goose Up
-- +goose StatementBegin
INSERT INTO car_photos(car_id, photo_url) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','https://picsum.photos/200/300' );
INSERT INTO car_photos(car_id, photo_url) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','https://picsum.photos/200/300' );
INSERT INTO car_photos(car_id, photo_url) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','https://picsum.photos/200/300' );
INSERT INTO car_photos(car_id, photo_url) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','https://picsum.photos/200/300' );
INSERT INTO car_photos(car_id, photo_url) VALUES('5d3bd16e-b185-44cc-ae48-76b86532cffe','https://picsum.photos/200/300' );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE car_photos CASCADE;
-- +goose StatementEnd
