CREATE TABLE IF NOT EXISTS users (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  first_name TEXT NOT NULL,
  middle_name TEXT,
  last_name TEXT NOT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  driver_license  TEXT NOT NULL,
  phone_number TEXT NOT NULL,
  account_number TEXT NOT NULL,
  bank_name TEXT NOT NULL,
  is_owner BOOLEAN DEFAULT false,
  is_renter BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS cars (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  owner_id UUID REFERENCES users(id) NOT NULL,
  make TEXT NOT NULL,
  model TEXT NOT NULL,
  year TEXT NOT NULL,
  license_plate TEXT UNIQUE NOT NULL,
  vin_number TEXT UNIQUE NOT NULL,
  transmission TRANSMISSION  NOT NULL,
  fuel_type FUEL_TYPE NOT NULL,
  mileage INT NOT NULL,
  location GEOMETRY(POINT, 4326) NOT NULL,
  price_per_hour NUMERIC(10,2) NOT NULL,
  status STATUS_TYPE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
