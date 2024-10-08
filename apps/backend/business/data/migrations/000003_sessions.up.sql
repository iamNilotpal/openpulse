CREATE TABLE
  user_sessions (
    user_id BIGINT NOT NULL REFERENCES users (id),
    session_token VARCHAR(255) NOT NULL UNIQUE,
    user_agent TEXT NOT NULL,
    ip_address TEXT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    device_info JSONB,
    location_info JSONB,
    revoked_at TIMESTAMP,
    last_activity_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP not NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id, session_token),
    INDEX (user_id),
    INDEX (user_id, is_active)
  );