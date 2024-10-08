CREATE TYPE system_appearance_type AS ENUM ('light', 'dark', 'system');

CREATE TYPE user_account_status_type AS ENUM ('active', 'suspended', 'deleted');

CREATE TABLE
  IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash BYTEA NOT NULL,
    phone_number VARCHAR(15) UNIQUE,
    avatar_url TEXT,
    is_verified BOOLEAN NOT NULL,
    account_status user_account_status_type DEFAULT 'active',
    organization_designation VARCHAR(100),
    role_id SMALLINT NOT NULL REFERENCES roles (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS users_access_controls (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users (id),
    role_id SMALLINT NOT NULL REFERENCES roles (id),
    resource_id SMALLINT NOT NULL REFERENCES resources (id),
    permission_id SMALLINT NOT NULL REFERENCES permissions (id),
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, role_id, resource_id, permission_id),
    INDEX (user_id, role_id)
  );

CREATE TABLE
  IF NOT EXISTS users_preferences (
    id BIGSERIAL PRIMARY KEY NOT NULl,
    user_id BIGINT NOT NULL REFERENCES users (id),
    appearance system_appearance_type DEFAULT 'system',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id)
  );

ALTER TABLE users
ADD COLUMN preference_id BIGINT REFERENCES users_preferences (id);