CREATE TYPE system_appearance_type AS ENUM ('light', 'dark', 'system');

CREATE TABLE
  IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash BYTEA,
    phone_number VARCHAR(15) UNIQUE,
    avatar_url TEXT,
    designation VARCHAR(100),
    account_status SMALLINT NOT NULL DEFAULT 1,
    is_email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    role_id SMALLINT NOT NULL REFERENCES roles (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS users_preferences (
    id BIGSERIAL PRIMARY KEY NOT NULl,
    user_id BIGINT NOT NULL REFERENCES users (id),
    appearance system_appearance_type NOT NULL DEFAULT 'system',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id)
  );

ALTER TABLE users
ADD COLUMN preference_id BIGINT REFERENCES users_preferences (id);

CREATE TABLE
  IF NOT EXISTS oauth_accounts (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    provider VARCHAR(50) NOT NULL,
    external_id VARCHAR(255) UNIQUE NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    scope TEXT,
    metadata JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (provider, user_id, external_id)
  );

CREATE INDEX idx_oauth_accounts_oauth_user_id ON oauth_accounts (user_id);

CREATE INDEX idx_oauth_accounts_oauth_provider ON oauth_accounts (provider);