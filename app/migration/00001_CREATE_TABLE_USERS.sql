-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id            UUID PRIMARY KEY gen_random_uuid(),
    email         VARCHAR(255) UNIQUE NOT NULL,
    phone         VARCHAR(20) UNIQUE,
    password_hash VARCHAR(255)        NOT NULL,
    first_name    VARCHAR(100)        NOT NULL,
    last_name     VARCHAR(100)        NOT NULL,
    avatar_url    VARCHAR(500),

    is_active     BOOLEAN   DEFAULT TRUE,
    is_verified   BOOLEAN   DEFAULT TRUE,
    last_login_at TIMESTAMP,

    created_at    TIMESTAMP DEFAULT NOW(),
    updated_at    TIMESTAMP DEFAULT NOW(),
    deleted_at    TIMESTAMP NULL,

    INDEX         idx_email (email),
    INDEX         idx_phone (phone),
    INDEX         idx_deleted (deleted_at)
);

-- +goose StatementEnd

-- +goose Down