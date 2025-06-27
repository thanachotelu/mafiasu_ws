create EXTENSION if not exists "uuid-ossp";
create type user_status as enum ('active','inactive','banned','pending');
create type car_status as enum ('active','inactive');
create type booking_status as enum ('confirmed','unconfirmed');

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

CREATE TABLE trackclicks (
    session_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_id uuid REFERENCES cars (car_id) NOT NULL,
    affiliator_id uuid REFERENCES affiliator (affiliator_id) NOT NULL,
    referral_link TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp
);

CREATE TABLE booking (
    book_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    session_id uuid REFERENCES trackclicks (session_id),
    user_id uuid REFERENCES users (user_id) NOT NULL,
    car_id uuid REFERENCES cars (car_id) NOT NULL,
    affiliator_id uuid REFERENCES affiliator (affiliator_id),
    total_price DECIMAL(10, 2) NOT NULL,
    pickup_date DATE,
    return_date DATE,
    status booking_status DEFAULT 'unconfirmed',
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp
);

CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    api_key TEXT UNIQUE NOT NULL
);

ALTER TABLE clients ADD COLUMN revoked BOOLEAN default false;

CREATE TABLE logs (
    id SERIAL PRIMARY KEY,
    client_id INT,
    user_id text,
    endpoint TEXT NOT NULL,
    method TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TABLE carimages (
    car_id uuid,
    image_path text,
    foreign key (car_id) REFERENCES cars(car_id)
);

INSERT INTO users (username, password_hash, firstname, lastname, phonenumber, email, role, status)
VALUES
(
    'thanachote',
    '$2a$12$Rm6efilT4b3OnTX/kVbgsujIehgJEfjNXme.VUWDJYoaa6qFTBBqu',
    'Thanachote',
    'Keakwanwong',
    '0891234567',
    'thanachote@example.com',
    'admin', 
    'active'
),
(
    'saturn',
    '$2a$12$ERh6a1Izt5py9iw3DEyFgeDE0CgdNidh801s3YOQ2aPSuoHUkWUdO',
    'Saturn',
    'Thanakul',
    '0907654321',
    'saturn@example.com',
    'user', 
    'active'
);

INSERT INTO cars (brand, model, license_plate, cartype, seat, doors, geartype, fueltype, rental_price_per_day, status, created_at, updated_at) VALUES
('Toyota', 'Hiace', '1 กย 5012', 'รถตู้', 12, 4, 'อัตโนมัติ', 'เบนซิน', 2080.00, 'active', NOW(), NOW()),
('Toyota', 'Commuter', '2กข 4821', 'รถตู้', 12, 4, 'อัตโนมัติ', 'เบนซิน', 2180.00, 'active', NOW(), NOW()),
('Hyundai', 'Staria', '4ขจ 1739', 'รถตู้', 11, 4, 'อัตโนมัติ', 'เบนซิน', 3180.00, 'active', NOW(), NOW()),
('Hyundai', 'New H1 รุ่น Elite เสริมประตูไฟฟ้า', '5งบ 9273', 'รถตู้', 7, 5, 'อัตโนมัติ', 'ดีเซล', 2500.00, 'active', NOW(), NOW()),
('Toyota', 'Fortuner', '1 กย 5012', 'SUV', 7, 4, 'อัตโนมัติ', 'เบนซิน', 1880.00, 'active', NOW(), NOW()),
('Toyota', 'Sienta', '3คด 0152', 'รถตู้', 7, 4, 'อัตโนมัติ', 'เบนซิน', 1320.00, 'active', NOW(), NOW()),
('Toyota', 'Avanza', '1จร 6048', 'รถตู้', 7, 4, 'อัตโนมัติ', 'เบนซิน', 1320.00, 'active', NOW(), NOW()),
('Toyota', 'Camry', '7ขท 8317', 'รถเก๋ง', 5, 4, 'อัตโนมัติ', 'เบนซิน', 1780.00, 'active', NOW(), NOW()),
('Honda', 'Accord', '2ณฟ 7921', 'รถเก๋ง', 5, 4, 'อัตโนมัติ', 'เบนซิน', 2200.00, 'active', NOW(), NOW()),
('Toyota', 'New Hybrid Camry', '9ดน 3106', 'รถเก๋ง', 5, 4, 'อัตโนมัติ', 'เบนซิน', 2180.00, 'active', NOW(), NOW()),
('Honda', 'Accord 1.5 TURBO', '6นย 5045', 'รถเก๋ง', 5, 4, 'อัตโนมัติ', 'เบนซิน', 1850.00, 'active', NOW(), NOW()),
('Isuzu', 'D-Max Cab4 1.9AT', '8พค 1698', 'กระบะ 4 ประตู', 5, 4, 'อัตโนมัติ CVT', 'ดีเซล', 1690.00, 'active', NOW(), NOW()),
('Mitsubishi', 'Xpander All New (2024)', '8ปพ 2874', 'เอนกประสงค์ MPV', 7, 5, 'อัตโนมัติ CVT', 'เบนซิน', 1590.00, 'active', NOW(), NOW()),
('Isuzu', 'All New Mu-X', '3บว 1360', 'SUV', 7, 5, 'อัตโนมัติ 6 สปีด', 'ดีเซล', 2300.00, 'active', NOW(), NOW()),
('Honda', 'City 1.0 Turbo (RS)', 'ศน 2483', 'รถยนต์ ECO-Car', 5, 4, 'อัตโนมัติ', 'เบนซิน', 1390.00, 'active', NOW(), NOW()),
('Honda', 'City 1.0 Turbo (RS)', '5พษ 9185', 'รถยนต์ ECO-Car', 5, 4, 'อัตโนมัติ', 'เบนซิน', 1390.00, 'active', NOW(), NOW()),
('Hyundai', 'New H1 รุ่น Elite เสริมประตูไฟฟ้า', '7สข 4871', 'รถตู้', 7, 5, 'อัตโนมัติ', 'ดีเซล', 2500.00, 'active', NOW(), NOW()),
('Toyota', 'Yaris', 'ภค 7012', 'รถยนต์ ECO-Car', 5, 4, 'อัตโนมัติ', 'เบนซิน', 780.00, 'active', NOW(), NOW()),
('Toyota', 'Vios', 'มช 5239', 'รถยนต์ ECO-Car', 5, 4, 'อัตโนมัติ', 'เบนซิน', 980.00, 'active', NOW(), NOW()),
('Toyota', 'Altis', '8กฟ 6524', 'รถยนต์ ECO-Car', 5, 4, 'อัตโนมัติ', 'เบนซิน', 1200.00, 'active', NOW(), NOW()),
('Toyota', 'Revo', '9ณว 2157', 'กระบะ 4 ประตู', 5, 4, 'อัตโนมัติ CVT', 'ดีเซล', 1200.00, 'active', NOW(), NOW()),
('Isuzu', 'D-Max', '3ณจ 5097', 'กระบะ 4 ประตู', 5, 4, 'อัตโนมัติ CVT', 'ดีเซล', 1250.00, 'active', NOW(), NOW()),
('Ford', 'Ranger', '7ขบ 1482', 'กระบะ 4 ประตู', 5, 4, 'อัตโนมัติ CVT', 'ดีเซล', 1100.00, 'active', NOW(), NOW());

insert into carimages (car_id, image_path) values 
((select car_id from cars where model = 'Hiace' and license_plate = '1 กย 5012' limit 1),'assets/images/Hiace.png'),
((select car_id from cars where model = 'Commuter' and license_plate = '2กข 4821' limit 1),'assets/images/Commuter.png'),
((select car_id from cars where model = 'Staria' and license_plate = '4ขจ 1739' limit 1),'assets/images/Staria.png'),
((select car_id from cars where model = 'New H1 รุ่น Elite เสริมประตูไฟฟ้า' and license_plate = '5งบ 9273' limit 1),'assets/images/New H1 รุ่น Elite เสริมประตูไฟฟ้า.png'),
((select car_id from cars where model = 'Fortuner' and license_plate = '1 กย 5012' limit 1),'assets/images/Fortuner.png'),
((select car_id from cars where model = 'Sienta' and license_plate = '3คด 0152' limit 1),'assets/images/Sienta.png'),
((select car_id from cars where model = 'Avanza' and license_plate = '1จร 6048' limit 1),'assets/images/Avanza.png'),
((select car_id from cars where model = 'Camry' and license_plate = '7ขท 8317' limit 1),'assets/images/Camry.png'),
((select car_id from cars where model = 'Accord' and license_plate = '2ณฟ 7921' limit 1),'assets/images/Accord.png'),
((select car_id from cars where model = 'New Hybrid Camry' and license_plate = '9ดน 3106' limit 1),'assets/images/New Hybrid Camry.png'),
((select car_id from cars where model = 'Accord 1.5 TURBO' and license_plate = '6นย 5045' limit 1),'assets/images/Accord 1.5 TURBO.png'),
((select car_id from cars where model = 'D-Max Cab4 1.9AT' and license_plate = '8พค 1698' limit 1),'assets/images/D-Max Cab4 1.9AT.png'),
((select car_id from cars where model = 'Xpander All New (2024)' and license_plate = '8ปพ 2874' limit 1),'assets/images/Xpander All New.png'),
((select car_id from cars where model = 'All New Mu-X' and license_plate = '3บว 1360' limit 1),'assets/images/All New Mu-X.png'),
((select car_id from cars where model = 'City 1.0 Turbo (RS)' and license_plate = 'ศน 2483' limit 1),'assets/images/City 1.0 Turbo (RS).png'),
((select car_id from cars where model = 'City 1.0 Turbo (RS)' and license_plate = '5พษ 9185' limit 1),'assets/images/City 1.0 Turbo (RS).png'),
((select car_id from cars where model = 'New H1 รุ่น Elite เสริมประตูไฟฟ้า' and license_plate = '7สข 4871'  limit 1),'assets/images/New H1 รุ่น Elite เสริมประตูไฟฟ้า.png'),
((select car_id from cars where model = 'Yaris' and license_plate = 'ภค 7012' limit 1),'assets/images/Yaris.png'),
((select car_id from cars where model = 'Vios' and license_plate = 'มช 5239' limit 1),'assets/images/Vios.png'),
((select car_id from cars where model = 'Altis' and license_plate = '8กฟ 6524' limit 1),'assets/images/Altis.png'),
((select car_id from cars where model = 'Revo' and license_plate = '9ณว 2157' limit 1),'assets/images/Revo.png'),
((select car_id from cars where model = 'D-Max' and license_plate = '3ณจ 5097' limit 1),'assets/images/D-Max.png'),
((select car_id from cars where model = 'Ranger' and license_plate = '7ขบ 1482' limit 1),'assets/images/Ranger.png');

insert into affiliator (user_id,affiliate_code,referral_link,commission_rate,total_commission,balance,created_at,updated_at) values
((select user_id from users where username = 'saturn'),'ABC1234','example@test.com',5,20000,20000,NOW(),NOW());

insert into booking (user_id,car_id,total_price,pickup_date,return_date,created_at,updated_at) values
((select user_id from users where username = 'saturn'),(select car_id from cars where model = 'Staria'),2500,'2025-05-01','2025-05-03',NOW(),NOW());
