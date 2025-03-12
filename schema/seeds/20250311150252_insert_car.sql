-- +goose Up
-- +goose StatementBegin
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('c72ff794-3d1d-4583-a988-eb89fe509ee3', 'Tesla', 'Model Y3', '2023', 'CE323A', '4B1SL65848Z411438', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable') ON CONFLICT (license_plate) DO NOTHING;
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('c72ff794-3d1d-4583-a988-eb89fe509ee3', 'Tesla', 'Model Y3', '2023', 'CE323A', '4B1SL65848Z411438', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable') ON CONFLICT (license_plate) DO NOTHING;
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('c72ff794-3d1d-4583-a988-eb89fe509ee3', 'Tesla', 'Model Y3', '2023', 'CE323A', '4B1SL65848Z411438', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable') ON CONFLICT (license_plate) DO NOTHING;
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('c72ff794-3d1d-4583-a988-eb89fe509ee3', 'Tesla', 'Model Y3', '2023', 'CE323A', '4B1SL65848Z411438', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable') ON CONFLICT (license_plate) DO NOTHING;
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('c72ff794-3d1d-4583-a988-eb89fe509ee3', 'Tesla', 'Model Y3', '2023', 'CE323A', '4B1SL65848Z411438', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable') ON CONFLICT (license_plate) DO NOTHING;
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('c72ff794-3d1d-4583-a988-eb89fe509ee3', 'Tesla', 'Model Y3', '2023', 'CE323A', '4B1SL65848Z411438', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable') ON CONFLICT (license_plate) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE cars CASCADE;
-- +goose StatementEnd
