CREATE TABLE
  IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash BYTEA NOT NULL,
    phone_number VARCHAR(15) UNIQUE,
    avatar_url TEXT,
    account_status VARCHAR(15) CHECK (
      account_status IN ('active', 'suspended', "deleted")
    ) DEFAULT 'active',
    role_id SMALLINT NOT NULL REFERENCES roles (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS users_preferences (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL UNIQUE,
    timezone VARCHAR(30),
    appearance VARCHAR(6) NOT NULL CHECK (theme IN ('light', 'dark', 'system')) DEFAULT 'system',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS users_permissions (
    user_id BIGINT NOT NULL REFERENCES users (id),
    permission_id SMALLINT NOT NULL REFERENCES permissions (id),
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by BIGINT NOT NULL REFERENCES users (id),
    UNIQUE (user_id, permission_id),
    INDEX user_id INCLUDE (permission_id)
  );