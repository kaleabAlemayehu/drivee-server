-- +goose Up
-- +goose StatementBegin
-- Seed data for the car_photos table
-- Assumes the car UUIDs below match the ones inserted into the cars table

INSERT INTO car_photos (car_id, photo_url) VALUES
-- Car 1 Photos (Toyota Camry)
('c0000001-1111-4444-8888-000000000001', 'http://localhost:9090/static/carphoto1.png'),
('c0000001-1111-4444-8888-000000000001', 'http://localhost:9090/static/carphoto2.png'),
('c0000001-1111-4444-8888-000000000001', 'http://localhost:9090/static/carphoto3.png'),
('c0000001-1111-4444-8888-000000000001', 'http://localhost:9090/static/carphoto4.png'),
('c0000001-1111-4444-8888-000000000001', 'http://localhost:9090/static/carphoto5.png'),

-- Car 2 Photos (Honda Civic)
('c0000002-1111-4444-8888-000000000002', 'http://localhost:9090/static/carphoto2.png'),
('c0000002-1111-4444-8888-000000000002', 'http://localhost:9090/static/carphoto6.png'),
('c0000002-1111-4444-8888-000000000002', 'http://localhost:9090/static/carphoto7.png'),
('c0000002-1111-4444-8888-000000000002', 'http://localhost:9090/static/carphoto8.png'),
('c0000002-1111-4444-8888-000000000002', 'http://localhost:9090/static/carphoto9.png'),
('c0000002-1111-4444-8888-000000000002', 'http://localhost:9090/static/carphoto10.png'),

-- Car 3 Photos (Ford F-150)
('c0000003-1111-4444-8888-000000000003', 'http://localhost:9090/static/carphoto3.png'),
('c0000003-1111-4444-8888-000000000003', 'http://localhost:9090/static/carphoto11.png'),
('c0000003-1111-4444-8888-000000000003', 'http://localhost:9090/static/carphoto12.png'),
('c0000003-1111-4444-8888-000000000003', 'http://localhost:9090/static/carphoto1.png'),
('c0000003-1111-4444-8888-000000000003', 'http://localhost:9090/static/carphoto2.png'),

-- Car 4 Photos (Chevrolet Malibu)
('c0000004-1111-4444-8888-000000000004', 'http://localhost:9090/static/carphoto4.png'),
('c0000004-1111-4444-8888-000000000004', 'http://localhost:9090/static/carphoto5.png'),
('c0000004-1111-4444-8888-000000000004', 'http://localhost:9090/static/carphoto6.png'),
('c0000004-1111-4444-8888-000000000004', 'http://localhost:9090/static/carphoto7.png'),
('c0000004-1111-4444-8888-000000000004', 'http://localhost:9090/static/carphoto8.png'),

-- Car 5 Photos (Nissan Altima)
('c0000005-1111-4444-8888-000000000005', 'http://localhost:9090/static/carphoto5.png'),
('c0000005-1111-4444-8888-000000000005', 'http://localhost:9090/static/carphoto9.png'),
('c0000005-1111-4444-8888-000000000005', 'http://localhost:9090/static/carphoto10.png'),
('c0000005-1111-4444-8888-000000000005', 'http://localhost:9090/static/carphoto11.png'),
('c0000005-1111-4444-8888-000000000005', 'http://localhost:9090/static/carphoto12.png'),

-- Car 6 Photos (BMW 3 Series)
('c0000006-1111-4444-8888-000000000006', 'http://localhost:9090/static/carphoto6.png'),
('c0000006-1111-4444-8888-000000000006', 'http://localhost:9090/static/carphoto1.png'),
('c0000006-1111-4444-8888-000000000006', 'http://localhost:9090/static/carphoto2.png'),
('c0000006-1111-4444-8888-000000000006', 'http://localhost:9090/static/carphoto3.png'),
('c0000006-1111-4444-8888-000000000006', 'http://localhost:9090/static/carphoto4.png'),

-- Car 7 Photos (Jeep Wrangler)
('c0000007-1111-4444-8888-000000000007', 'http://localhost:9090/static/carphoto7.png'),
('c0000007-1111-4444-8888-000000000007', 'http://localhost:9090/static/carphoto8.png'),
('c0000007-1111-4444-8888-000000000007', 'http://localhost:9090/static/carphoto9.png'),
('c0000007-1111-4444-8888-000000000007', 'http://localhost:9090/static/carphoto10.png'),
('c0000007-1111-4444-8888-000000000007', 'http://localhost:9090/static/carphoto11.png'),
('c0000007-1111-4444-8888-000000000007', 'http://localhost:9090/static/carphoto12.png'),

-- Car 8 Photos (Tesla Model 3)
('c0000008-1111-4444-8888-000000000008', 'http://localhost:9090/static/carphoto8.png'),
('c0000008-1111-4444-8888-000000000008', 'http://localhost:9090/static/carphoto1.png'),
('c0000008-1111-4444-8888-000000000008', 'http://localhost:9090/static/carphoto3.png'),
('c0000008-1111-4444-8888-000000000008', 'http://localhost:9090/static/carphoto5.png'),
('c0000008-1111-4444-8888-000000000008', 'http://localhost:9090/static/carphoto7.png'),

-- Car 9 Photos (Subaru Outback)
('c0000009-1111-4444-8888-000000000009', 'http://localhost:9090/static/carphoto9.png'),
('c0000009-1111-4444-8888-000000000009', 'http://localhost:9090/static/carphoto2.png'),
('c0000009-1111-4444-8888-000000000009', 'http://localhost:9090/static/carphoto4.png'),
('c0000009-1111-4444-8888-000000000009', 'http://localhost:9090/static/carphoto6.png'),
('c0000009-1111-4444-8888-000000000009', 'http://localhost:9090/static/carphoto8.png'),

-- Car 10 Photos (Hyundai Sonata)
('c0000010-1111-4444-8888-000000000010', 'http://localhost:9090/static/carphoto10.png'),
('c0000010-1111-4444-8888-000000000010', 'http://localhost:9090/static/carphoto11.png'),
('c0000010-1111-4444-8888-000000000010', 'http://localhost:9090/static/carphoto12.png'),
('c0000010-1111-4444-8888-000000000010', 'http://localhost:9090/static/carphoto1.png'),
('c0000010-1111-4444-8888-000000000010', 'http://localhost:9090/static/carphoto3.png'),

-- Car 11 Photos (Kia Sportage)
('c0000011-1111-4444-8888-000000000011', 'http://localhost:9090/static/carphoto11.png'),
('c0000011-1111-4444-8888-000000000011', 'http://localhost:9090/static/carphoto5.png'),
('c0000011-1111-4444-8888-000000000011', 'http://localhost:9090/static/carphoto6.png'),
('c0000011-1111-4444-8888-000000000011', 'http://localhost:9090/static/carphoto7.png'),
('c0000011-1111-4444-8888-000000000011', 'http://localhost:9090/static/carphoto9.png'),

-- Car 12 Photos (Mazda CX-5)
('c0000012-1111-4444-8888-000000000012', 'http://localhost:9090/static/carphoto12.png'),
('c0000012-1111-4444-8888-000000000012', 'http://localhost:9090/static/carphoto2.png'),
('c0000012-1111-4444-8888-000000000012', 'http://localhost:9090/static/carphoto4.png'),
('c0000012-1111-4444-8888-000000000012', 'http://localhost:9090/static/carphoto8.png'),
('c0000012-1111-4444-8888-000000000012', 'http://localhost:9090/static/carphoto10.png'),

-- Car 13 Photos (Volkswagen Jetta)
('c0000013-1111-4444-8888-000000000013', 'http://localhost:9090/static/carphoto1.png'),
('c0000013-1111-4444-8888-000000000013', 'http://localhost:9090/static/carphoto3.png'),
('c0000013-1111-4444-8888-000000000013', 'http://localhost:9090/static/carphoto5.png'),
('c0000013-1111-4444-8888-000000000013', 'http://localhost:9090/static/carphoto7.png'),
('c0000013-1111-4444-8888-000000000013', 'http://localhost:9090/static/carphoto9.png'),

-- Car 14 Photos (Toyota RAV4)
('c0000014-1111-4444-8888-000000000014', 'http://localhost:9090/static/carphoto2.png'),
('c0000014-1111-4444-8888-000000000014', 'http://localhost:9090/static/carphoto4.png'),
('c0000014-1111-4444-8888-000000000014', 'http://localhost:9090/static/carphoto6.png'),
('c0000014-1111-4444-8888-000000000014', 'http://localhost:9090/static/carphoto8.png'),
('c0000014-1111-4444-8888-000000000014', 'http://localhost:9090/static/carphoto10.png'),
('c0000014-1111-4444-8888-000000000014', 'http://localhost:9090/static/carphoto12.png'),

-- Car 15 Photos (Ford Escape)
('c0000015-1111-4444-8888-000000000015', 'http://localhost:9090/static/carphoto3.png'),
('c0000015-1111-4444-8888-000000000015', 'http://localhost:9090/static/carphoto5.png'),
('c0000015-1111-4444-8888-000000000015', 'http://localhost:9090/static/carphoto7.png'),
('c0000015-1111-4444-8888-000000000015', 'http://localhost:9090/static/carphoto9.png'),
('c0000015-1111-4444-8888-000000000015', 'http://localhost:9090/static/carphoto11.png'),

-- Car 16 Photos (Honda CR-V)
('c0000016-1111-4444-8888-000000000016', 'http://localhost:9090/static/carphoto4.png'),
('c0000016-1111-4444-8888-000000000016', 'http://localhost:9090/static/carphoto6.png'),
('c0000016-1111-4444-8888-000000000016', 'http://localhost:9090/static/carphoto8.png'),
('c0000016-1111-4444-8888-000000000016', 'http://localhost:9090/static/carphoto10.png'),
('c0000016-1111-4444-8888-000000000016', 'http://localhost:9090/static/carphoto12.png'),

-- Car 17 Photos (Chevrolet Equinox)
('c0000017-1111-4444-8888-000000000017', 'http://localhost:9090/static/carphoto5.png'),
('c0000017-1111-4444-8888-000000000017', 'http://localhost:9090/static/carphoto1.png'),
('c0000017-1111-4444-8888-000000000017', 'http://localhost:9090/static/carphoto3.png'),
('c0000017-1111-4444-8888-000000000017', 'http://localhost:9090/static/carphoto7.png'),
('c0000017-1111-4444-8888-000000000017', 'http://localhost:9090/static/carphoto9.png'),
('c0000017-1111-4444-8888-000000000017', 'http://localhost:9090/static/carphoto11.png'),

-- Car 18 Photos (Nissan Rogue)
('c0000018-1111-4444-8888-000000000018', 'http://localhost:9090/static/carphoto6.png'),
('c0000018-1111-4444-8888-000000000018', 'http://localhost:9090/static/carphoto2.png'),
('c0000018-1111-4444-8888-000000000018', 'http://localhost:9090/static/carphoto4.png'),
('c0000018-1111-4444-8888-000000000018', 'http://localhost:9090/static/carphoto8.png'),
('c0000018-1111-4444-8888-000000000018', 'http://localhost:9090/static/carphoto10.png'),

-- Car 19 Photos (BMW X3)
('c0000019-1111-4444-8888-000000000019', 'http://localhost:9090/static/carphoto7.png'),
('c0000019-1111-4444-8888-000000000019', 'http://localhost:9090/static/carphoto1.png'),
('c0000019-1111-4444-8888-000000000019', 'http://localhost:9090/static/carphoto3.png'),
('c0000019-1111-4444-8888-000000000019', 'http://localhost:9090/static/carphoto5.png'),
('c0000019-1111-4444-8888-000000000019', 'http://localhost:9090/static/carphoto9.png'),

-- Car 20 Photos (Jeep Grand Cherokee)
('c0000020-1111-4444-8888-000000000020', 'http://localhost:9090/static/carphoto8.png'),
('c0000020-1111-4444-8888-000000000020', 'http://localhost:9090/static/carphoto2.png'),
('c0000020-1111-4444-8888-000000000020', 'http://localhost:9090/static/carphoto4.png'),
('c0000020-1111-4444-8888-000000000020', 'http://localhost:9090/static/carphoto6.png'),
('c0000020-1111-4444-8888-000000000020', 'http://localhost:9090/static/carphoto10.png'),
('c0000020-1111-4444-8888-000000000020', 'http://localhost:9090/static/carphoto12.png'),

-- Car 21 Photos (Tesla Model Y)
('c0000021-1111-4444-8888-000000000021', 'http://localhost:9090/static/carphoto9.png'),
('c0000021-1111-4444-8888-000000000021', 'http://localhost:9090/static/carphoto11.png'),
('c0000021-1111-4444-8888-000000000021', 'http://localhost:9090/static/carphoto1.png'),
('c0000021-1111-4444-8888-000000000021', 'http://localhost:9090/static/carphoto3.png'),
('c0000021-1111-4444-8888-000000000021', 'http://localhost:9090/static/carphoto5.png'),

-- Car 22 Photos (Subaru Forester)
('c0000022-1111-4444-8888-000000000022', 'http://localhost:9090/static/carphoto10.png'),
('c0000022-1111-4444-8888-000000000022', 'http://localhost:9090/static/carphoto12.png'),
('c0000022-1111-4444-8888-000000000022', 'http://localhost:9090/static/carphoto2.png'),
('c0000022-1111-4444-8888-000000000022', 'http://localhost:9090/static/carphoto4.png'),
('c0000022-1111-4444-8888-000000000022', 'http://localhost:9090/static/carphoto6.png'),

-- Car 23 Photos (Hyundai Tucson)
('c0000023-1111-4444-8888-000000000023', 'http://localhost:9090/static/carphoto11.png'),
('c0000023-1111-4444-8888-000000000023', 'http://localhost:9090/static/carphoto1.png'),
('c0000023-1111-4444-8888-000000000023', 'http://localhost:9090/static/carphoto3.png'),
('c0000023-1111-4444-8888-000000000023', 'http://localhost:9090/static/carphoto5.png'),
('c0000023-1111-4444-8888-000000000023', 'http://localhost:9090/static/carphoto7.png'),
('c0000023-1111-4444-8888-000000000023', 'http://localhost:9090/static/carphoto9.png'),

-- Car 24 Photos (Kia Sorento)
('c0000024-1111-4444-8888-000000000024', 'http://localhost:9090/static/carphoto12.png'),
('c0000024-1111-4444-8888-000000000024', 'http://localhost:9090/static/carphoto2.png'),
('c0000024-1111-4444-8888-000000000024', 'http://localhost:9090/static/carphoto4.png'),
('c0000024-1111-4444-8888-000000000024', 'http://localhost:9090/static/carphoto6.png'),
('c0000024-1111-4444-8888-000000000024', 'http://localhost:9090/static/carphoto8.png'),

-- Car 25 Photos (Mazda 3)
('c0000025-1111-4444-8888-000000000025', 'http://localhost:9090/static/carphoto1.png'),
('c0000025-1111-4444-8888-000000000025', 'http://localhost:9090/static/carphoto3.png'),
('c0000025-1111-4444-8888-000000000025', 'http://localhost:9090/static/carphoto5.png'),
('c0000025-1111-4444-8888-000000000025', 'http://localhost:9090/static/carphoto7.png'),
('c0000025-1111-4444-8888-000000000025', 'http://localhost:9090/static/carphoto9.png'),

-- Car 26 Photos (Volkswagen Tiguan)
('c0000026-1111-4444-8888-000000000026', 'http://localhost:9090/static/carphoto2.png'),
('c0000026-1111-4444-8888-000000000026', 'http://localhost:9090/static/carphoto4.png'),
('c0000026-1111-4444-8888-000000000026', 'http://localhost:9090/static/carphoto6.png'),
('c0000026-1111-4444-8888-000000000026', 'http://localhost:9090/static/carphoto8.png'),
('c0000026-1111-4444-8888-000000000026', 'http://localhost:9090/static/carphoto10.png'),
('c0000026-1111-4444-8888-000000000026', 'http://localhost:9090/static/carphoto12.png'),

-- Car 27 Photos (Toyota Highlander)
('c0000027-1111-4444-8888-000000000027', 'http://localhost:9090/static/carphoto3.png'),
('c0000027-1111-4444-8888-000000000027', 'http://localhost:9090/static/carphoto5.png'),
('c0000027-1111-4444-8888-000000000027', 'http://localhost:9090/static/carphoto7.png'),
('c0000027-1111-4444-8888-000000000027', 'http://localhost:9090/static/carphoto9.png'),
('c0000027-1111-4444-8888-000000000027', 'http://localhost:9090/static/carphoto11.png'),

-- Car 28 Photos (Ford Mustang)
('c0000028-1111-4444-8888-000000000028', 'http://localhost:9090/static/carphoto4.png'),
('c0000028-1111-4444-8888-000000000028', 'http://localhost:9090/static/carphoto6.png'),
('c0000028-1111-4444-8888-000000000028', 'http://localhost:9090/static/carphoto8.png'),
('c0000028-1111-4444-8888-000000000028', 'http://localhost:9090/static/carphoto10.png'),
('c0000028-1111-4444-8888-000000000028', 'http://localhost:9090/static/carphoto12.png'),

-- Car 29 Photos (Honda Accord)
('c0000029-1111-4444-8888-000000000029', 'http://localhost:9090/static/carphoto5.png'),
('c0000029-1111-4444-8888-000000000029', 'http://localhost:9090/static/carphoto1.png'),
('c0000029-1111-4444-8888-000000000029', 'http://localhost:9090/static/carphoto3.png'),
('c0000029-1111-4444-8888-000000000029', 'http://localhost:9090/static/carphoto7.png'),
('c0000029-1111-4444-8888-000000000029', 'http://localhost:9090/static/carphoto9.png'),
('c0000029-1111-4444-8888-000000000029', 'http://localhost:9090/static/carphoto11.png'),

-- Car 30 Photos (Chevrolet Traverse)
('c0000030-1111-4444-8888-000000000030', 'http://localhost:9090/static/carphoto6.png'),
('c0000030-1111-4444-8888-000000000030', 'http://localhost:9090/static/carphoto2.png'),
('c0000030-1111-4444-8888-000000000030', 'http://localhost:9090/static/carphoto4.png'),
('c0000030-1111-4444-8888-000000000030', 'http://localhost:9090/static/carphoto8.png'),
('c0000030-1111-4444-8888-000000000030', 'http://localhost:9090/static/carphoto10.png');

-- Seed data for car_photos for the 10 additional cars

INSERT INTO car_photos (car_id, photo_url) VALUES
-- Car 31 Photos (Audi A4)
('c0000031-1111-4444-8888-000000000031', 'http://localhost:9090/static/carphoto7.png'),
('c0000031-1111-4444-8888-000000000031', 'http://localhost:9090/static/carphoto8.png'),
('c0000031-1111-4444-8888-000000000031', 'http://localhost:9090/static/carphoto9.png'),
('c0000031-1111-4444-8888-000000000031', 'http://localhost:9090/static/carphoto10.png'),
('c0000031-1111-4444-8888-000000000031', 'http://localhost:9090/static/carphoto11.png'),

-- Car 32 Photos (Lexus RX)
('c0000032-1111-4444-8888-000000000032', 'http://localhost:9090/static/carphoto8.png'),
('c0000032-1111-4444-8888-000000000032', 'http://localhost:9090/static/carphoto12.png'),
('c0000032-1111-4444-8888-000000000032', 'http://localhost:9090/static/carphoto1.png'),
('c0000032-1111-4444-8888-000000000032', 'http://localhost:9090/static/carphoto2.png'),
('c0000032-1111-4444-8888-000000000032', 'http://localhost:9090/static/carphoto3.png'),
('c0000032-1111-4444-8888-000000000032', 'http://localhost:9090/static/carphoto4.png'),

-- Car 33 Photos (Ford Explorer)
('c0000033-1111-4444-8888-000000000033', 'http://localhost:9090/static/carphoto9.png'),
('c0000033-1111-4444-8888-000000000033', 'http://localhost:9090/static/carphoto5.png'),
('c0000033-1111-4444-8888-000000000033', 'http://localhost:9090/static/carphoto6.png'),
('c0000033-1111-4444-8888-000000000033', 'http://localhost:9090/static/carphoto7.png'),
('c0000033-1111-4444-8888-000000000033', 'http://localhost:9090/static/carphoto8.png'),

-- Car 34 Photos (Toyota Prius)
('c0000034-1111-4444-8888-000000000034', 'http://localhost:9090/static/carphoto10.png'),
('c0000034-1111-4444-8888-000000000034', 'http://localhost:9090/static/carphoto11.png'),
('c0000034-1111-4444-8888-000000000034', 'http://localhost:9090/static/carphoto12.png'),
('c0000034-1111-4444-8888-000000000034', 'http://localhost:9090/static/carphoto1.png'),
('c0000034-1111-4444-8888-000000000034', 'http://localhost:9090/static/carphoto2.png'),

-- Car 35 Photos (Chevrolet Silverado)
('c0000035-1111-4444-8888-000000000035', 'http://localhost:9090/static/carphoto11.png'),
('c0000035-1111-4444-8888-000000000035', 'http://localhost:9090/static/carphoto3.png'),
('c0000035-1111-4444-8888-000000000035', 'http://localhost:9090/static/carphoto4.png'),
('c0000035-1111-4444-8888-000000000035', 'http://localhost:9090/static/carphoto5.png'),
('c0000035-1111-4444-8888-000000000035', 'http://localhost:9090/static/carphoto6.png'),
('c0000035-1111-4444-8888-000000000035', 'http://localhost:9090/static/carphoto7.png'),

-- Car 36 Photos (Mini Cooper)
('c0000036-1111-4444-8888-000000000036', 'http://localhost:9090/static/carphoto12.png'),
('c0000036-1111-4444-8888-000000000036', 'http://localhost:9090/static/carphoto8.png'),
('c0000036-1111-4444-8888-000000000036', 'http://localhost:9090/static/carphoto9.png'),
('c0000036-1111-4444-8888-000000000036', 'http://localhost:9090/static/carphoto10.png'),
('c0000036-1111-4444-8888-000000000036', 'http://localhost:9090/static/carphoto11.png'),

-- Car 37 Photos (Hyundai Palisade)
('c0000037-1111-4444-8888-000000000037', 'http://localhost:9090/static/carphoto1.png'),
('c0000037-1111-4444-8888-000000000037', 'http://localhost:9090/static/carphoto2.png'),
('c0000037-1111-4444-8888-000000000037', 'http://localhost:9090/static/carphoto3.png'),
('c0000037-1111-4444-8888-000000000037', 'http://localhost:9090/static/carphoto4.png'),
('c0000037-1111-4444-8888-000000000037', 'http://localhost:9090/static/carphoto5.png'),

-- Car 38 Photos (Ram 1500)
('c0000038-1111-4444-8888-000000000038', 'http://localhost:9090/static/carphoto2.png'),
('c0000038-1111-4444-8888-000000000038', 'http://localhost:9090/static/carphoto6.png'),
('c0000038-1111-4444-8888-000000000038', 'http://localhost:9090/static/carphoto7.png'),
('c0000038-1111-4444-8888-000000000038', 'http://localhost:9090/static/carphoto8.png'),
('c0000038-1111-4444-8888-000000000038', 'http://localhost:9090/static/carphoto9.png'),
('c0000038-1111-4444-8888-000000000038', 'http://localhost:9090/static/carphoto10.png'),

-- Car 39 Photos (Kia Telluride)
('c0000039-1111-4444-8888-000000000039', 'http://localhost:9090/static/carphoto3.png'),
('c0000039-1111-4444-8888-000000000039', 'http://localhost:9090/static/carphoto11.png'),
('c0000039-1111-4444-8888-000000000039', 'http://localhost:9090/static/carphoto12.png'),
('c0000039-1111-4444-8888-000000000039', 'http://localhost:9090/static/carphoto1.png'),
('c0000039-1111-4444-8888-000000000039', 'http://localhost:9090/static/carphoto2.png'),

-- Car 40 Photos (Porsche Macan)
('c0000040-1111-4444-8888-000000000040', 'http://localhost:9090/static/carphoto4.png'),
('c0000040-1111-4444-8888-000000000040', 'http://localhost:9090/static/carphoto5.png'),
('c0000040-1111-4444-8888-000000000040', 'http://localhost:9090/static/carphoto6.png'),
('c0000040-1111-4444-8888-000000000040', 'http://localhost:9090/static/carphoto7.png'),
('c0000040-1111-4444-8888-000000000040', 'http://localhost:9090/static/carphoto8.png');




-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE car_photos CASCADE;
-- +goose StatementEnd
