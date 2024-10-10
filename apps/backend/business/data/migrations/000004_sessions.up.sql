CREATE TABLE
  user_sessions (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users (id),
    session_token VARCHAR(255) NOT NULL UNIQUE,
    user_agent TEXT NOT NULL,
    ip_address TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    device_info JSONB,
    location_info JSONB,
    revoked_at TIMESTAMP,
    expires_at BIGINT NOT NULL,
    created_at TIMESTAMP not NULL DEFAULT CURRENT_TIMESTAMP,
    last_activity_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX (user_id),
    INDEX (is_active)
  );