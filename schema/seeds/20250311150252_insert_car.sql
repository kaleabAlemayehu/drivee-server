-- +goose Up
-- +goose StatementBegin
INSERT INTO cars (
  owner_id, make, model, year, thumbnail_picture, license_plate, vin_number,
  transmission, fuel_type, mileage, location, price_per_hour, status
) VALUES
-- 1
(uuid_generate_v4(), 'Toyota', 'Camry', '2018', 'https://example.com/camry.jpg',
'ABC1234', '1HGCM82633A004352', 'automatic', 'gas', 45200,
ST_PointFromText('POINT(-118.2437 34.0522)', 4326), 12.50, 'avaliable'),
-- 2
(uuid_generate_v4(), 'Honda', 'Civic', '2020', 'https://example.com/civic.jpg',
'XYZ5678', '2HGFB2F50DH513698', 'automatic', 'gas', 22000,
ST_PointFromText('POINT(-122.4194 37.7749)', 4326), 14.00, 'avaliable'),
-- 3
(uuid_generate_v4(), 'Tesla', 'Model 3', '2022', 'https://example.com/model3.jpg',
'TESLA99', '5YJ3E1EA7KF317829', 'automatic', 'electric', 9000,
ST_PointFromText('POINT(-73.935242 40.730610)', 4326), 25.00, 'avaliable'),
-- 4
(uuid_generate_v4(), 'Ford', 'F-150', '2019', 'https://example.com/f150.jpg',
'FORD789', '1FTEW1EP1KKD51234', 'automatic', 'gas', 38000,
ST_PointFromText('POINT(-80.191790 25.761680)', 4326), 18.00, 'avaliable'),
-- 5
(uuid_generate_v4(), 'Chevrolet', 'Malibu', '2017', 'https://example.com/malibu.jpg',
'CHEV456', '1G1ZD5ST3HF123456', 'automatic', 'gas', 61000,
ST_PointFromText('POINT(-87.6298 41.8781)', 4326), 10.00, 'avaliable'),
-- 6
(uuid_generate_v4(), 'BMW', 'M3', '2016', 'https://example.com/m3.jpg',
'BMWM333', 'WBS3C9C56FP123456', 'manual', 'gas', 52000,
ST_PointFromText('POINT(-0.1276 51.5072)', 4326), 35.00, 'avaliable'),
-- 7
(uuid_generate_v4(), 'Subaru', 'Outback', '2015', 'https://example.com/outback.jpg',
'SUB9999', '4S4BSANC1G1234567', 'automatic', 'gas', 71000,
ST_PointFromText('POINT(144.9631 -37.8136)', 4326), 11.00, 'avaliable'),
-- 8
(uuid_generate_v4(), 'Volkswagen', 'Golf', '2018', 'https://example.com/golf.jpg',
'VWGOLF1', '3VWD07AJ5EM123456', 'manual', 'diesel', 47000,
ST_PointFromText('POINT(2.3522 48.8566)', 4326), 13.50, 'avaliable'),
-- 9
(uuid_generate_v4(), 'Audi', 'A4', '2021', 'https://example.com/a4.jpg',
'AUDI444', 'WAUENAF48MN123456', 'automatic', 'gas', 16000,
ST_PointFromText('POINT(2.3522 48.8566)', 4326), 29.00, 'avaliable'),
-- 10
(uuid_generate_v4(), 'Kia', 'Soul', '2019', 'https://example.com/soul.jpg',
'KIAS123', 'KNDJN2A29K7654321', 'automatic', 'gas', 41000,
ST_PointFromText('POINT(139.6917 35.6895)', 4326), 13.00, 'avaliable'),
-- 11
(uuid_generate_v4(), 'Mazda', 'MX-5', '2017', 'https://example.com/mx5.jpg',
'MAZDA77', 'JM1NDAD7XG0101010', 'manual', 'gas', 34000,
ST_PointFromText('POINT(-3.7038 40.4168)', 4326), 20.00, 'avaliable'),
-- 12
(uuid_generate_v4(), 'Nissan', 'Leaf', '2022', 'https://example.com/leaf.jpg',
'NISSLEAF', '1N4AZ1CP5KC123456', 'automatic', 'electric', 8000,
ST_PointFromText('POINT(13.4050 52.5200)', 4326), 19.00, 'avaliable'),
-- 13
(uuid_generate_v4(), 'Hyundai', 'Ioniq', '2021', 'https://example.com/ioniq.jpg',
'HYUNION', 'KMHC75LH1MU123456', 'automatic', 'hybrid', 14000,
ST_PointFromText('POINT(151.2093 -33.8688)', 4326), 18.00, 'avaliable'),
-- 14
(uuid_generate_v4(), 'Ford', 'Mustang', '2018', 'https://example.com/mustang.jpg',
'MUSTNG66', '1FA6P8TH8J1234567', 'manual', 'gas', 33000,
ST_PointFromText('POINT(-84.3880 33.7490)', 4326), 22.00, 'avaliable'),
-- 15
(uuid_generate_v4(), 'Toyota', 'Prius', '2020', 'https://example.com/prius.jpg',
'PRIUS00', 'JTDKARFU0L3123456', 'automatic', 'hybrid', 24000,
ST_PointFromText('POINT(-74.0060 40.7128)', 4326), 16.00, 'avaliable'),
-- 16
(uuid_generate_v4(), 'Chevrolet', 'Bolt', '2021', 'https://example.com/bolt.jpg',
'BOLT321', '1G1FY6S00M1234567', 'automatic', 'electric', 13000,
ST_PointFromText('POINT(-97.7431 30.2672)', 4326), 21.00, 'avaliable'),
-- 17
(uuid_generate_v4(), 'Honda', 'Accord', '2019', 'https://example.com/accord.jpg',
'HONACC99', '1HGCV1F34KA123456', 'automatic', 'gas', 38000,
ST_PointFromText('POINT(-95.3698 29.7604)', 4326), 15.00, 'avaliable'),
-- 18
(uuid_generate_v4(), 'Jeep', 'Wrangler', '2017', 'https://example.com/wrangler.jpg',
'JEEP555', '1C4HJXDN1HW123456', 'manual', 'gas', 62000,
ST_PointFromText('POINT(-112.0740 33.4484)', 4326), 17.00, 'avaliable'),
-- 19
(uuid_generate_v4(), 'Dodge', 'Charger', '2018', 'https://example.com/charger.jpg',
'DODGE88', '2C3CDXCT9JH123456', 'automatic', 'gas', 44000,
ST_PointFromText('POINT(-122.3321 47.6062)', 4326), 20.00, 'avaliable'),
-- 20
(uuid_generate_v4(), 'BMW', 'i3', '2020', 'https://example.com/i3.jpg',
'BMWI3333', 'WBY8P4C5XL7D12345', 'automatic', 'electric', 17000,
ST_PointFromText('POINT(174.7633 -36.8485)', 4326), 26.00, 'avaliable'),
-- 21
(uuid_generate_v4(), 'Mercedes-Benz', 'E-Class', '2022', 'https://example.com/eclass.jpg',
'MERC333', 'WDDZF8DB3LA123456', 'automatic', 'diesel', 7000,
ST_PointFromText('POINT(12.4964 41.9028)', 4326), 40.00, 'avaliable'),
-- 22
(uuid_generate_v4(), 'Volvo', 'XC90', '2021', 'https://example.com/xc90.jpg',
'VOLVO90', 'YV4A22PK7M1234567', 'automatic', 'hybrid', 15000,
ST_PointFromText('POINT(18.0686 59.3293)', 4326), 34.00, 'avaliable'),
-- 23
(uuid_generate_v4(), 'Peugeot', '208', '2019', 'https://example.com/208.jpg',
'PEU2080', 'VR3UPHN0KE1234567', 'manual', 'diesel', 39000,
ST_PointFromText('POINT(2.3522 48.8566)', 4326), 12.00, 'avaliable'),
-- 24
(uuid_generate_v4(), 'Renault', 'Zoe', '2022', 'https://example.com/zoe.jpg',
'RENAU11', 'VF1AGVYB45234567', 'automatic', 'electric', 5000,
ST_PointFromText('POINT(2.2137 46.2276)', 4326), 19.00, 'avaliable'),
-- 25
(uuid_generate_v4(), 'Fiat', '500e', '2021', 'https://example.com/500e.jpg',
'FIAT500E', '3C3CFFGE7MT123456', 'automatic', 'electric', 11000,
ST_PointFromText('POINT(12.4964 41.9028)', 4326), 17.00, 'avaliable'),
-- 26
(uuid_generate_v4(), 'Mini', 'Cooper', '2017', 'https://example.com/cooper.jpg',
'MINI123', 'WMWXM5C53HT123456', 'manual', 'gas', 43000,
ST_PointFromText('POINT(0.1276 51.5072)', 4326), 16.00, 'avaliable'),
-- 27
(uuid_generate_v4(), 'Subaru', 'Forester', '2018', 'https://example.com/forester.jpg',
'SUBFO33', 'JF2SJAGC5JH123456', 'automatic', 'gas', 37000,
ST_PointFromText('POINT(135.7681 35.0116)', 4326), 14.00, 'avaliable'),
-- 28
(uuid_generate_v4(), 'Hyundai', 'Kona', '2020', 'https://example.com/kona.jpg',
'KONA001', 'KM8K1CAA0LU123456', 'automatic', 'electric', 19000,
ST_PointFromText('POINT(126.9780 37.5665)', 4326), 20.00, 'avaliable'),
-- 29
(uuid_generate_v4(), 'Ford', 'Fusion', '2016', 'https://example.com/fusion.jpg',
'FUSION1', '3FA6P0H71GR123456', 'automatic', 'hybrid', 66000,
ST_PointFromText('POINT(-122.6765 45.5231)', 4326), 12.00, 'avaliable'),
-- 30
(uuid_generate_v4(), 'Lexus', 'RX 350', '2021', 'https://example.com/rx350.jpg',
'LEXRX350', '2T2HZMDA9MC123456', 'automatic', 'gas', 12000,
ST_PointFromText('POINT(-71.0589 42.3601)', 4326), 32.00, 'avaliable');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE cars CASCADE;
-- +goose StatementEnd
