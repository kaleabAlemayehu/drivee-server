-- +goose Up
-- +goose StatementBegin
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('30746294-1228-480b-a2df-3fcadd83f337', 'Tesla', 'Model Y3', '2023', 'CA323A', '4Y1SL65848Z411439', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable');
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('30746294-1228-480b-a2df-3fcadd83f337', 'Tesla', 'Model Y3', '2023', 'CB323A', '4Y1SL65848Z411439', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable');
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('30746294-1228-480b-a2df-3fcadd83f337', 'Tesla', 'Model Y3', '2023', 'CC323A', '4Y1SL65848Z411439', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable');
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('30746294-1228-480b-a2df-3fcadd83f337', 'Tesla', 'Model Y3', '2023', 'CD323A', '4Y1SL65848Z411439', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable');
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('30746294-1228-480b-a2df-3fcadd83f337', 'Tesla', 'Model Y3', '2023', 'CE323A', '4Y1SL65848Z411439', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable');
INSERT INTO cars (owner_id, make, model, year, license_plate, vin_number , transmission, fuel_type, mileage, location, price_per_hour, status ) VALUES ('30746294-1228-480b-a2df-3fcadd83f337', 'Tesla', 'Model Y3', '2023', 'CF323A', '4Y1SL65848Z411439', 'automatic', 'hybrid', 232,ST_SetSRID(ST_MakePoint(-71.060316, 48.432044), 4326),99, 'avaliable');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE cars CASCADE;
-- +goose StatementEnd
