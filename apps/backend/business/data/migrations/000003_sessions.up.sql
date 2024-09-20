CREATE TABLE
  user_sessions (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT REFERENCES users (id),
    session_token VARCHAR(255) NOT NULL UNIQUE,
    refresh_token VARCHAR(255) NOT NULL UNIQUE,
    user_agent TEXT,
    ip_address VARCHAR(45),
    created_at TIMESTAMP
    WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      expires_at TIMESTAMP
    WITH
      TIME ZONE,
      last_activity_at TIMESTAMP
    WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      is_active BOOLEAN DEFAULT TRUE,
      device_info JSONB,
      location_info JSONB,
      revoked_at TIMESTAMP
    WITH
      TIME ZONE,
      INDEX (user_id),
      INDEX (is_active),
      INDEX (session_token),
      INDEX (refresh_token)
  );