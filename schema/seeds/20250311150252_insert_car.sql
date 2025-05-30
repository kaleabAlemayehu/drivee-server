-- +goose Up
-- +goose StatementBegin
-- Seed data for the cars table
-- Using ENUM types:
-- CREATE TYPE STATUS_TYPE AS ENUM('avaliable', 'rented', 'inactive');
-- CREATE TYPE TRANSMISSION AS ENUM('manual', 'automatic');
-- CREATE TYPE FUEL_TYPE AS ENUM('gas', 'diesel', 'electric', 'hybrid');
-- Assuming PostGIS extension is enabled for GEOMETRY type
-- Assuming uuid-ossp extension is enabled for uuid_generate_v4()

-- Define some owner UUIDs (replace with actual user IDs later)
DO $$
DECLARE
    owner1_id UUID := '95ce777e-50a5-4138-b895-6b799070dd3b';
    owner2_id UUID := '95ce777e-50a5-4138-b895-6b799070dd3b';
    owner3_id UUID := '95ce777e-50a5-4138-b895-6b799070dd3b';
    owner4_id UUID := '95ce777e-50a5-4138-b895-6b799070dd3b';
    owner5_id UUID := '95ce777e-50a5-4138-b895-6b799070dd3b';
    owner6_id UUID := '95ce777e-50a5-4138-b895-6b799070dd3b';
    owner7_id UUID := '4441605e-87e3-4916-adc0-3f79a2d29d29';
    owner8_id UUID := '4441605e-87e3-4916-adc0-3f79a2d29d29';
    owner9_id UUID := '4441605e-87e3-4916-adc0-3f79a2d29d29';
    owner10_id UUID := '4441605e-87e3-4916-adc0-3f79a2d29d29';
BEGIN
    -- Insert Cars (Generating specific UUIDs for linking photos)
    INSERT INTO cars (id, owner_id, thumbnail_picture, description, make, model, year, license_plate, vin_number, transmission, fuel_type, mileage, location, price_per_hour, status) VALUES
    ('c0000001-1111-4444-8888-000000000001', owner1_id, 'http://localhost:9090/static/carphoto1.png', 'Reliable Toyota Camry, great for city trips.', 'Toyota', 'Camry', '2021', 'CAR001', 'VIN000000000000001', 'automatic', 'gas', 25000, ST_SetSRID(ST_MakePoint(-74.0060, 40.7128), 4326), 15.50, 'avaliable'),
    ('c0000002-1111-4444-8888-000000000002', owner2_id, 'http://localhost:9090/static/carphoto2.png', 'Honda Civic, fuel efficient and fun to drive.', 'Honda', 'Civic', '2020', 'CAR002', 'VIN000000000000002', 'automatic', 'gas', 32000, ST_SetSRID(ST_MakePoint(-73.9857, 40.7484), 4326), 14.00, 'avaliable'),
    ('c0000003-1111-4444-8888-000000000003', owner3_id, 'http://localhost:9090/static/carphoto3.png', 'Ford F-150, powerful truck for work or play.', 'Ford', 'F-150', '2019', 'CAR003', 'VIN000000000000003', 'automatic', 'gas', 45000, ST_SetSRID(ST_MakePoint(-73.9654, 40.7829), 4326), 22.00, 'rented'),
    ('c0000004-1111-4444-8888-000000000004', owner1_id, 'http://localhost:9090/static/carphoto4.png', 'Chevrolet Malibu, comfortable mid-size sedan.', 'Chevrolet', 'Malibu', '2022', 'CAR004', 'VIN000000000000004', 'automatic', 'gas', 15000, ST_SetSRID(ST_MakePoint(-74.0020, 40.7230), 4326), 16.00, 'avaliable'),
    ('c0000005-1111-4444-8888-000000000005', owner4_id, 'http://localhost:9090/static/carphoto5.png', 'Nissan Altima, sleek design and smooth ride.', 'Nissan', 'Altima', '2021', 'CAR005', 'VIN000000000000005', 'automatic', 'gas', 28000, ST_SetSRID(ST_MakePoint(-73.9900, 40.7500), 4326), 15.00, 'inactive'),
    ('c0000006-1111-4444-8888-000000000006', owner5_id, 'http://localhost:9090/static/carphoto6.png', 'BMW 3 Series, luxury and performance combined.', 'BMW', '3 Series', '2020', 'CAR006', 'VIN000000000000006', 'automatic', 'gas', 38000, ST_SetSRID(ST_MakePoint(-73.9750, 40.7600), 4326), 25.50, 'avaliable'),
    ('c0000007-1111-4444-8888-000000000007', owner2_id, 'http://localhost:9090/static/carphoto7.png', 'Jeep Wrangler, ultimate off-road capability.', 'Jeep', 'Wrangler', '2018', 'CAR007', 'VIN000000000000007', 'manual', 'gas', 55000, ST_SetSRID(ST_MakePoint(-73.9550, 40.7700), 4326), 20.00, 'avaliable'),
    ('c0000008-1111-4444-8888-000000000008', owner3_id, 'http://localhost:9090/static/carphoto8.png', 'Tesla Model 3, electric innovation.', 'Tesla', 'Model 3', '2022', 'CAR008', 'VIN000000000000008', 'automatic', 'electric', 12000, ST_SetSRID(ST_MakePoint(-74.0100, 40.7050), 4326), 30.00, 'rented'),
    ('c0000009-1111-4444-8888-000000000009', owner4_id, 'http://localhost:9090/static/carphoto9.png', 'Subaru Outback, reliable AWD for all conditions.', 'Subaru', 'Outback', '2019', 'CAR009', 'VIN000000000000009', 'automatic', 'gas', 48000, ST_SetSRID(ST_MakePoint(-73.9800, 40.7550), 4326), 18.50, 'avaliable'),
    ('c0000010-1111-4444-8888-000000000010', owner5_id, 'http://localhost:9090/static/carphoto10.png', 'Hyundai Sonata, stylish and feature-packed.', 'Hyundai', 'Sonata', '2021', 'CAR010', 'VIN000000000000010', 'automatic', 'gas', 22000, ST_SetSRID(ST_MakePoint(-73.9700, 40.7650), 4326), 14.50, 'avaliable'),
    ('c0000011-1111-4444-8888-000000000011', owner1_id, 'http://localhost:9090/static/carphoto11.png', 'Kia Sportage, compact SUV with great value.', 'Kia', 'Sportage', '2020', 'CAR011', 'VIN000000000000011', 'automatic', 'gas', 34000, ST_SetSRID(ST_MakePoint(-73.9600, 40.7750), 4326), 17.00, 'inactive'),
    ('c0000012-1111-4444-8888-000000000012', owner2_id, 'http://localhost:9090/static/carphoto12.png', 'Mazda CX-5, sporty handling in an SUV.', 'Mazda', 'CX-5', '2022', 'CAR012', 'VIN000000000000012', 'automatic', 'gas', 18000, ST_SetSRID(ST_MakePoint(-73.9500, 40.7800), 4326), 19.00, 'avaliable'),
    ('c0000013-1111-4444-8888-000000000013', owner3_id, 'http://localhost:9090/static/carphoto1.png', 'Volkswagen Jetta, German engineering, practical.', 'Volkswagen', 'Jetta', '2019', 'CAR013', 'VIN000000000000013', 'automatic', 'gas', 41000, ST_SetSRID(ST_MakePoint(-74.0050, 40.7150), 4326), 13.50, 'avaliable'),
    ('c0000014-1111-4444-8888-000000000014', owner4_id, 'http://localhost:9090/static/carphoto2.png', 'Toyota RAV4, popular and versatile SUV.', 'Toyota', 'RAV4', '2021', 'CAR014', 'VIN000000000000014', 'automatic', 'hybrid', 26000, ST_SetSRID(ST_MakePoint(-73.9880, 40.7420), 4326), 18.00, 'rented'),
    ('c0000015-1111-4444-8888-000000000015', owner5_id, 'http://localhost:9090/static/carphoto3.png', 'Ford Escape, compact and efficient SUV.', 'Ford', 'Escape', '2020', 'CAR015', 'VIN000000000000015', 'automatic', 'gas', 33000, ST_SetSRID(ST_MakePoint(-73.9720, 40.7880), 4326), 16.50, 'avaliable'),
    ('c0000016-1111-4444-8888-000000000016', owner1_id, 'http://localhost:9090/static/carphoto4.png', 'Honda CR-V, spacious and reliable family SUV.', 'Honda', 'CR-V', '2022', 'CAR016', 'VIN000000000000016', 'automatic', 'gas', 14000, ST_SetSRID(ST_MakePoint(-74.0010, 40.7280), 4326), 19.50, 'avaliable'),
    ('c0000017-1111-4444-8888-000000000017', owner2_id, 'http://localhost:9090/static/carphoto5.png', 'Chevrolet Equinox, comfortable and tech-savvy.', 'Chevrolet', 'Equinox', '2019', 'CAR017', 'VIN000000000000017', 'automatic', 'gas', 49000, ST_SetSRID(ST_MakePoint(-73.9950, 40.7520), 4326), 17.50, 'inactive'),
    ('c0000018-1111-4444-8888-000000000018', owner3_id, 'http://localhost:9090/static/carphoto6.png', 'Nissan Rogue, popular choice for families.', 'Nissan', 'Rogue', '2021', 'CAR018', 'VIN000000000000018', 'automatic', 'gas', 29000, ST_SetSRID(ST_MakePoint(-73.9780, 40.7620), 4326), 18.00, 'avaliable'),
    ('c0000019-1111-4444-8888-000000000019', owner4_id, 'http://localhost:9090/static/carphoto7.png', 'BMW X3, luxury compact SUV experience.', 'BMW', 'X3', '2020', 'CAR019', 'VIN000000000000019', 'automatic', 'gas', 39000, ST_SetSRID(ST_MakePoint(-73.9620, 40.7720), 4326), 28.00, 'rented'),
    ('c0000020-1111-4444-8888-000000000020', owner5_id, 'http://localhost:9090/static/carphoto8.png', 'Jeep Grand Cherokee, capable and refined SUV.', 'Jeep', 'Grand Cherokee', '2018', 'CAR020', 'VIN000000000000020', 'automatic', 'gas', 60000, ST_SetSRID(ST_MakePoint(-73.9480, 40.7820), 4326), 23.00, 'avaliable'),
    ('c0000021-1111-4444-8888-000000000021', owner1_id, 'http://localhost:9090/static/carphoto9.png', 'Tesla Model Y, versatile electric SUV.', 'Tesla', 'Model Y', '2022', 'CAR021', 'VIN000000000000021', 'automatic', 'electric', 11000, ST_SetSRID(ST_MakePoint(-74.0120, 40.7080), 4326), 32.00, 'avaliable'),
    ('c0000022-1111-4444-8888-000000000022', owner2_id, 'http://localhost:9090/static/carphoto10.png', 'Subaru Forester, practical and safe AWD.', 'Subaru', 'Forester', '2019', 'CAR022', 'VIN000000000000022', 'automatic', 'gas', 51000, ST_SetSRID(ST_MakePoint(-73.9820, 40.7580), 4326), 19.00, 'avaliable'),
    ('c0000023-1111-4444-8888-000000000023', owner3_id, 'http://localhost:9090/static/carphoto11.png', 'Hyundai Tucson, modern design, good features.', 'Hyundai', 'Tucson', '2021', 'CAR023', 'VIN000000000000023', 'automatic', 'gas', 24000, ST_SetSRID(ST_MakePoint(-73.9710, 40.7680), 4326), 17.00, 'rented'),
    ('c0000024-1111-4444-8888-000000000024', owner4_id, 'http://localhost:9090/static/carphoto12.png', 'Kia Sorento, spacious 3-row SUV option.', 'Kia', 'Sorento', '2020', 'CAR024', 'VIN000000000000024', 'automatic', 'gas', 36000, ST_SetSRID(ST_MakePoint(-73.9610, 40.7780), 4326), 21.00, 'avaliable'),
    ('c0000025-1111-4444-8888-000000000025', owner5_id, 'http://localhost:9090/static/carphoto1.png', 'Mazda 3, fun-to-drive compact car.', 'Mazda', '3', '2022', 'CAR025', 'VIN000000000000025', 'manual', 'gas', 16000, ST_SetSRID(ST_MakePoint(-73.9510, 40.7810), 4326), 15.00, 'avaliable'),
    ('c0000026-1111-4444-8888-000000000026', owner1_id, 'http://localhost:9090/static/carphoto2.png', 'Volkswagen Tiguan, roomy compact SUV.', 'Volkswagen', 'Tiguan', '2019', 'CAR026', 'VIN000000000000026', 'automatic', 'gas', 43000, ST_SetSRID(ST_MakePoint(-74.0030, 40.7180), 4326), 18.50, 'inactive'),
    ('c0000027-1111-4444-8888-000000000027', owner2_id, 'http://localhost:9090/static/carphoto3.png', 'Toyota Highlander, reliable 3-row SUV.', 'Toyota', 'Highlander', '2021', 'CAR027', 'VIN000000000000027', 'automatic', 'hybrid', 27000, ST_SetSRID(ST_MakePoint(-73.9860, 40.7450), 4326), 24.00, 'avaliable'),
    ('c0000028-1111-4444-8888-000000000028', owner3_id, 'http://localhost:9090/static/carphoto4.png', 'Ford Mustang, iconic American muscle car.', 'Ford', 'Mustang', '2020', 'CAR028', 'VIN000000000000028', 'manual', 'gas', 31000, ST_SetSRID(ST_MakePoint(-73.9680, 40.7850), 4326), 26.00, 'rented'),
    ('c0000029-1111-4444-8888-000000000029', owner4_id, 'http://localhost:9090/static/carphoto5.png', 'Honda Accord, benchmark mid-size sedan.', 'Honda', 'Accord', '2022', 'CAR029', 'VIN000000000000029', 'automatic', 'gas', 13000, ST_SetSRID(ST_MakePoint(-74.0000, 40.7250), 4326), 17.00, 'avaliable'),
    ('c0000030-1111-4444-8888-000000000030', owner5_id, 'http://localhost:9090/static/carphoto6.png', 'Chevrolet Traverse, spacious family hauler.', 'Chevrolet', 'Traverse', '2019', 'CAR030', 'VIN000000000000030', 'automatic', 'gas', 52000, ST_SetSRID(ST_MakePoint(-73.9920, 40.7510), 4326), 20.50, 'avaliable');
END $$;

-- Seed data for 10 additional cars for a specific owner
-- Using ENUM types: STATUS_TYPE, TRANSMISSION, FUEL_TYPE
-- Assuming PostGIS and uuid-ossp extensions are enabled

DO $$
DECLARE
    specific_owner_id UUID := 'f8e5eb15-5991-41a5-ae50-68005a77d9c7';
BEGIN
    -- Insert 10 more Cars for the specified owner
    INSERT INTO cars (id, owner_id, thumbnail_picture, description, make, model, year, license_plate, vin_number, transmission, fuel_type, mileage, location, price_per_hour, status) VALUES
    ('c0000031-1111-4444-8888-000000000031', specific_owner_id, 'http://localhost:9090/static/carphoto7.png', 'Audi A4, premium sedan with smooth handling.', 'Audi', 'A4', '2020', 'CAR031', 'VIN000000000000031', 'automatic', 'gas', 35000, ST_SetSRID(ST_MakePoint(-73.9580, 40.7680), 4326), 27.00, 'avaliable'),
    ('c0000032-1111-4444-8888-000000000032', specific_owner_id, 'http://localhost:9090/static/carphoto8.png', 'Lexus RX, luxury crossover with comfort.', 'Lexus', 'RX', '2019', 'CAR032', 'VIN000000000000032', 'automatic', 'hybrid', 42000, ST_SetSRID(ST_MakePoint(-73.9790, 40.7590), 4326), 29.50, 'rented'),
    ('c0000033-1111-4444-8888-000000000033', specific_owner_id, 'http://localhost:9090/static/carphoto9.png', 'Ford Explorer, spacious SUV for family adventures.', 'Ford', 'Explorer', '2021', 'CAR033', 'VIN000000000000033', 'automatic', 'gas', 28000, ST_SetSRID(ST_MakePoint(-74.0070, 40.7110), 4326), 23.50, 'avaliable'),
    ('c0000034-1111-4444-8888-000000000034', specific_owner_id, 'http://localhost:9090/static/carphoto10.png', 'Toyota Prius, the classic fuel-efficient hybrid.', 'Toyota', 'Prius', '2022', 'CAR034', 'VIN000000000000034', 'automatic', 'hybrid', 10000, ST_SetSRID(ST_MakePoint(-73.9890, 40.7410), 4326), 16.00, 'avaliable'),
    ('c0000035-1111-4444-8888-000000000035', specific_owner_id, 'http://localhost:9090/static/carphoto11.png', 'Chevrolet Silverado, heavy-duty pickup truck.', 'Chevrolet', 'Silverado', '2018', 'CAR035', 'VIN000000000000035', 'automatic', 'diesel', 65000, ST_SetSRID(ST_MakePoint(-73.9690, 40.7810), 4326), 24.50, 'inactive'),
    ('c0000036-1111-4444-8888-000000000036', specific_owner_id, 'http://localhost:9090/static/carphoto12.png', 'Mini Cooper, iconic style, fun to drive.', 'Mini', 'Cooper', '2021', 'CAR036', 'VIN000000000000036', 'manual', 'gas', 19000, ST_SetSRID(ST_MakePoint(-74.0090, 40.7210), 4326), 18.00, 'avaliable'),
    ('c0000037-1111-4444-8888-000000000037', specific_owner_id, 'http://localhost:9090/static/carphoto1.png', 'Hyundai Palisade, large SUV with premium features.', 'Hyundai', 'Palisade', '2022', 'CAR037', 'VIN000000000000037', 'automatic', 'gas', 15000, ST_SetSRID(ST_MakePoint(-73.9990, 40.7510), 4326), 25.00, 'rented'),
    ('c0000038-1111-4444-8888-000000000038', specific_owner_id, 'http://localhost:9090/static/carphoto2.png', 'Ram 1500, capable and comfortable truck.', 'Ram', '1500', '2020', 'CAR038', 'VIN000000000000038', 'automatic', 'gas', 40000, ST_SetSRID(ST_MakePoint(-73.9740, 40.7610), 4326), 23.00, 'avaliable'),
    ('c0000039-1111-4444-8888-000000000039', specific_owner_id, 'http://localhost:9090/static/carphoto3.png', 'Kia Telluride, award-winning family SUV.', 'Kia', 'Telluride', '2021', 'CAR039', 'VIN000000000000039', 'automatic', 'gas', 22000, ST_SetSRID(ST_MakePoint(-73.9540, 40.7710), 4326), 24.00, 'avaliable'),
    ('c0000040-1111-4444-8888-000000000040', specific_owner_id, 'http://localhost:9090/static/carphoto4.png', 'Porsche Macan, sporty luxury compact SUV.', 'Porsche', 'Macan', '2019', 'CAR040', 'VIN000000000000040', 'automatic', 'gas', 48000, ST_SetSRID(ST_MakePoint(-73.9440, 40.7810), 4326), 35.00, 'avaliable');
END $$;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE cars CASCADE;
-- +goose StatementEnd
