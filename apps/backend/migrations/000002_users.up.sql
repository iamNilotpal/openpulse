CREATE TYPE system_appearance_type AS ENUM ('light', 'dark', 'system');

CREATE TABLE
  IF NOT EXISTS users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    avatar_url TEXT,
    designation VARCHAR(100),
    country_code VARCHAR(5),
    phone_number VARCHAR(15),
    is_email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    account_status SMALLINT NOT NULL DEFAULT 1 CHECK (account_status > 0),
    role_id SMALLINT NOT NULL REFERENCES roles (id) ON DELETE NO ACTION,
    account_deleted_at BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (country_code, phone_number)
  );

CREATE INDEX idx_users_account_deleted_at ON users (account_deleted_at);

CREATE TABLE
  IF NOT EXISTS users_preferences (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULl,
    user_id BIGINT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    timezone VARCHAR(50) NOT NULL,
    appearance system_appearance_type NOT NULL DEFAULT 'system',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id)
  );

ALTER TABLE users
ADD COLUMN preference_id BIGINT REFERENCES users_preferences (id) ON DELETE SET NULL;

CREATE TABLE
  IF NOT EXISTS oauth_accounts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
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