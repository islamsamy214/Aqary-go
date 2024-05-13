-- name: CreateUser :one
INSERT INTO users (email, password, phone_number, profile_id)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetAllUsers :many
SELECT u.id, u.email, u.phone_number, p.first_name, p.last_name, p.address
FROM users u
JOIN profile p ON u.profile_id = p.id
ORDER BY p.first_name;

-- name: GenerateOTP :exec
UPDATE users
SET otp = $1, otp_expiration_time = $2
WHERE phone_number = $3;

-- name: VerifyOTP :one
SELECT otp_expiration_time
FROM users
WHERE phone_number = $1 AND otp = $2;

-- name: CreateProfile :one
INSERT INTO profile (first_name, last_name, address)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetProfileByID :one
SELECT first_name, last_name, address
FROM profile
WHERE id = $1;

-- name: UpdateProfileByID :exec
UPDATE profile
SET first_name = $2, last_name = $3, address = $4
WHERE id = $1;

-- name: DeleteProfileByID :exec
DELETE FROM profile
WHERE id = $1;