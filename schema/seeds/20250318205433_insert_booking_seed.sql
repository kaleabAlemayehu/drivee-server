-- +goose Up
-- +goose StatementBegin
-- Seed data for the bookings table
-- Referencing cars with IDs c0000031-... to c0000040-...
-- Assuming BOOKING_STATUS ENUM type exists
-- Assuming uuid-ossp extension is enabled

DO $$
DECLARE
    -- Define some Renter UUIDs (replace with actual user IDs)
    renter1_id UUID := 'f47ac10b-58cc-4372-a567-0e02b2c3d479';
    renter2_id UUID := 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d';
    renter3_id UUID := '123e4567-e89b-42d3-a456-556642440000';
    renter4_id UUID := '987e6543-e21b-48d7-a012-345642661111';
    -- Car IDs from the previous batch
    car31_id UUID := 'c0000031-1111-4444-8888-000000000031'; -- Audi A4, 27.00/hr
    car32_id UUID := 'c0000032-1111-4444-8888-000000000032'; -- Lexus RX, 29.50/hr
    car33_id UUID := 'c0000033-1111-4444-8888-000000000033'; -- Ford Explorer, 23.50/hr
    car34_id UUID := 'c0000034-1111-4444-8888-000000000034'; -- Toyota Prius, 16.00/hr
    car35_id UUID := 'c0000035-1111-4444-8888-000000000035'; -- Chevy Silverado, 24.50/hr
    car36_id UUID := 'c0000036-1111-4444-8888-000000000036'; -- Mini Cooper, 18.00/hr
    car37_id UUID := 'c0000037-1111-4444-8888-000000000037'; -- Hyundai Palisade, 25.00/hr
    car38_id UUID := 'c0000038-1111-4444-8888-000000000038'; -- Ram 1500, 23.00/hr
    car39_id UUID := 'c0000039-1111-4444-8888-000000000039'; -- Kia Telluride, 24.00/hr
    car40_id UUID := 'c0000040-1111-4444-8888-000000000040'; -- Porsche Macan, 35.00/hr
BEGIN
    -- Insert Bookings
    INSERT INTO bookings (car_id, renter_id, start_time, end_time, total_price, status) VALUES
    -- Completed Bookings (Past)
    (car31_id, renter1_id, '2025-04-10 10:00:00', '2025-04-12 10:00:00', (48 * 27.00), 'completed'), -- 2 days
    (car36_id, renter2_id, '2025-04-25 12:00:00', '2025-04-26 18:00:00', (30 * 18.00), 'completed'), -- 1.25 days
    (car39_id, renter1_id, '2025-04-15 14:00:00', '2025-04-18 14:00:00', (72 * 24.00), 'completed'), -- 3 days
    (car31_id, renter3_id, '2025-04-28 09:00:00', '2025-04-28 13:00:00', (4 * 27.00), 'completed'),   -- 4 hours
    (car38_id, renter2_id, '2025-04-22 07:00:00', '2025-04-24 19:00:00', (60 * 23.00), 'completed'), -- 2.5 days

    -- Confirmed Bookings (Ongoing or Future)
    (car32_id, renter2_id, '2025-05-05 09:00:00', '2025-05-05 17:00:00', (8 * 29.50), 'confirmed'),   -- Future (8 hours)
    (car35_id, renter4_id, '2025-05-01 15:00:00', '2025-05-04 15:00:00', (72 * 24.50), 'confirmed'), -- Ongoing (3 days)
    (car37_id, renter3_id, '2025-05-06 10:00:00', '2025-05-08 10:00:00', (48 * 25.00), 'confirmed'), -- Future (2 days)
    (car33_id, renter4_id, '2025-05-09 17:00:00', '2025-05-11 17:00:00', (48 * 23.50), 'confirmed'), -- Future (Weekend)
    (car40_id, renter3_id, '2025-05-12 14:00:00', '2025-05-14 10:00:00', (44 * 35.00), 'confirmed'), -- Future (~2 days)

    -- Pending Bookings (Future)
    (car33_id, renter1_id, '2025-05-10 12:00:00', '2025-05-11 12:00:00', (24 * 23.50), 'pending'),   -- Future (1 day)
    (car38_id, renter4_id, '2025-05-15 09:00:00', '2025-05-15 19:00:00', (10 * 23.00), 'pending'),   -- Future (10 hours)
    (car36_id, renter1_id, '2025-05-20 10:00:00', '2025-05-27 10:00:00', (168 * 18.00), 'pending'),  -- Future (1 week)

    -- Canceled Bookings
    (car34_id, renter3_id, '2025-04-20 08:00:00', '2025-04-20 18:00:00', (10 * 16.00), 'canceled'),  -- Past booking canceled
    (car40_id, renter2_id, '2025-05-07 11:00:00', '2025-05-09 11:00:00', (48 * 35.00), 'canceled'); -- Future booking canceled

END $$;
-- Additional seed data for the bookings table
-- Referencing cars with IDs c0000031-... to c0000040-...
-- Assuming BOOKING_STATUS ENUM type exists
-- Assuming uuid-ossp extension is enabled

DO $$
DECLARE
    -- Define Renter UUIDs (replace with actual user IDs)
    renter1_id UUID := 'f47ac10b-58cc-4372-a567-0e02b2c3d479';
    renter2_id UUID := 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d';
    renter3_id UUID := '123e4567-e89b-42d3-a456-556642440000';
    renter4_id UUID := '987e6543-e21b-48d7-a012-345642661111';
    renter5_id UUID := 'c8a9b0d1-e2f3-4a5b-8c7d-6e5f4a3b2c1e';

    -- Car IDs from the previous batch (and their prices for calculation)
    car31_id UUID := 'c0000031-1111-4444-8888-000000000031'; -- Audi A4, 27.00/hr
    car32_id UUID := 'c0000032-1111-4444-8888-000000000032'; -- Lexus RX, 29.50/hr
    car33_id UUID := 'c0000033-1111-4444-8888-000000000033'; -- Ford Explorer, 23.50/hr
    car34_id UUID := 'c0000034-1111-4444-8888-000000000034'; -- Toyota Prius, 16.00/hr
    car35_id UUID := 'c0000035-1111-4444-8888-000000000035'; -- Chevy Silverado, 24.50/hr
    car36_id UUID := 'c0000036-1111-4444-8888-000000000036'; -- Mini Cooper, 18.00/hr
    car37_id UUID := 'c0000037-1111-4444-8888-000000000037'; -- Hyundai Palisade, 25.00/hr
    car38_id UUID := 'c0000038-1111-4444-8888-000000000038'; -- Ram 1500, 23.00/hr
    car39_id UUID := 'c0000039-1111-4444-8888-000000000039'; -- Kia Telluride, 24.00/hr
    car40_id UUID := 'c0000040-1111-4444-8888-000000000040'; -- Porsche Macan, 35.00/hr
BEGIN
    -- Insert More Bookings
    INSERT INTO bookings (car_id, renter_id, start_time, end_time, total_price, status) VALUES
    -- More Completed Bookings (Past)
    (car32_id, renter3_id, '2025-04-05 10:00:00', '2025-04-07 10:00:00', (48 * 29.50), 'completed'), -- 2 days
    (car40_id, renter1_id, '2025-04-18 12:00:00', '2025-04-20 12:00:00', (48 * 35.00), 'completed'), -- 2 days
    (car34_id, renter4_id, '2025-04-29 09:00:00', '2025-04-30 17:00:00', (32 * 16.00), 'completed'), -- 1 day + 8 hours
    (car37_id, renter2_id, '2025-04-21 11:00:00', '2025-04-23 11:00:00', (48 * 25.00), 'completed'), -- 2 days
    (car39_id, renter5_id, '2025-04-26 10:00:00', '2025-04-26 20:00:00', (10 * 24.00), 'completed'), -- 10 hours

    -- More Confirmed Bookings (Future)
    (car31_id, renter2_id, '2025-05-08 10:00:00', '2025-05-10 10:00:00', (48 * 27.00), 'confirmed'), -- 2 days
    (car36_id, renter4_id, '2025-05-13 14:00:00', '2025-05-16 14:00:00', (72 * 18.00), 'confirmed'), -- 3 days
    (car38_id, renter1_id, '2025-05-18 08:00:00', '2025-05-19 18:00:00', (34 * 23.00), 'confirmed'), -- 1 day + 10 hours
    (car32_id, renter5_id, '2025-05-22 12:00:00', '2025-05-24 12:00:00', (48 * 29.50), 'confirmed'), -- 2 days
    (car34_id, renter1_id, '2025-05-07 09:00:00', '2025-05-07 17:00:00', (8 * 16.00), 'confirmed'),   -- 8 hours

    -- More Pending Bookings (Future)
    (car35_id, renter2_id, '2025-06-01 10:00:00', '2025-06-05 10:00:00', (96 * 24.50), 'pending'),   -- 4 days
    (car40_id, renter4_id, '2025-06-10 12:00:00', '2025-06-12 12:00:00', (48 * 35.00), 'pending'),   -- 2 days
    (car31_id, renter5_id, '2025-05-25 15:00:00', '2025-05-26 15:00:00', (24 * 27.00), 'pending'),   -- 1 day
    (car37_id, renter3_id, '2025-06-03 11:00:00', '2025-06-04 11:00:00', (24 * 25.00), 'pending'),   -- 1 day
    (car39_id, renter2_id, '2025-05-28 09:00:00', '2025-05-30 17:00:00', (56 * 24.00), 'pending'),   -- 2 days + 8 hours

    -- More Canceled Bookings
    (car33_id, renter2_id, '2025-05-15 10:00:00', '2025-05-17 10:00:00', (48 * 23.50), 'canceled'),  -- Future booking canceled
    (car36_id, renter5_id, '2025-05-03 13:00:00', '2025-05-03 21:00:00', (8 * 18.00), 'canceled');   -- Future booking canceled (short notice)

END $$;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE bookings CASCADE;
-- +goose StatementEnd
