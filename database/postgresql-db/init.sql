create EXTENSION if not exists "uuid-ossp";
create type user_status as enum ('active','inactive','banned','pending');
create type car_status as enum ('active','inactive');

create table users (
    user_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL,
    password_hash TEXT NOT NULL,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    phonenumber VARCHAR(12) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    role VARCHAR(50) default 'user',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    status user_status DEFAULT 'active'
);

CREATE TABLE affiliator (
    affiliator_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    affiliate_code VARCHAR(20) UNIQUE NOT NULL,
    referral_link TEXT,
    commission_rate DECIMAL(5,2) DEFAULT 10.00,
    total_commission DECIMAL(12,2) DEFAULT 0.00,
    balance DECIMAL(12,2) DEFAULT 0.00,
    api_key TEXT UNIQUE NOT NULL, -- เก็บ API Key ของ Affiliator
    created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION fn_create_affiliator()
RETURNS trigger AS $$
DECLARE
  code TEXT := substring(uuid_generate_v4()::text,1,8);
  key  TEXT := uuid_generate_v4()::text;
BEGIN
  INSERT INTO affiliator (
    affiliator_id, user_id, affiliate_code, referral_link, api_key
  ) VALUES (
    uuid_generate_v4(),
    NEW.user_id,
    code,
    'https://yourdomain.com/ref/' || code,
    key
  );
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


DROP TRIGGER IF EXISTS trg_after_user_insert ON users;
CREATE TRIGGER trg_after_user_insert
AFTER INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION fn_create_affiliator();

create table cars (
    car_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    brand VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    license_plate VARCHAR(50) NOT NULL,
    carType VARCHAR(50) NOT NULL,
    seat INT NOT NULL,
    doors INT NOT NULL,
    gearType VARCHAR(50),
    fuelType VARCHAR(50),
    rental_price_per_day DECIMAL(10, 2) NOT NULL,
    status car_status DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
);

create table booking (
    book_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid REFERENCES users (user_id) NOT NULL ,
    car_id uuid REFERENCES cars (car_id) NOT NULL,
    affiliator_id uuid REFERENCES affiliator (affiliator_id),
    total_price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
);


CREATE TABLE logs (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    endpoint TEXT NOT NULL,
    method TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

