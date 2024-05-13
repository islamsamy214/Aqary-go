-- Table: users
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    otp VARCHAR(6),
    otp_expiration_time TIMESTAMP,
    profile_id INT,
    FOREIGN KEY (profile_id) REFERENCES profile(id)
);

-- Table: profile
CREATE TABLE profile (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    address TEXT
);