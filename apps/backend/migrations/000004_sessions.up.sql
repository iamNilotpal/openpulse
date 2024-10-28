CREATE TABLE
  user_sessions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    session_token TEXT NOT NULL UNIQUE,
    user_agent TEXT NOT NULL,
    ip_address TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    device_info JSONB,
    location_info JSONB,
    revoked_at TIMESTAMP,
    expires_at BIGINT NOT NULL,
    created_at TIMESTAMP not NULL DEFAULT CURRENT_TIMESTAMP,
    last_activity_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

CREATE INDEX "idx_user_sessions_user_id" ON user_sessions (user_id);

CREATE INDEX "idx_user_sessions_is_active" ON user_sessions (is_active)