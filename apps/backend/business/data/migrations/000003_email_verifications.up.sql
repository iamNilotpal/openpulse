CREATE TABLE
  email_verifications (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT REFERENCES users (id) ON DELETE CASCADE,
    verification_token VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL,
    attempt_count SMALLINT NOT NULL DEFAULT 0,
    max_attempts SMALLINT NOT NULL DEFAULT 5,
    verified_at TIMESTAMP,
    expires_at BIGINT NOT NULL,
    is_revoked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX (user_id, expires_at, is_revoked)
  );