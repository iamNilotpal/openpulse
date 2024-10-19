CREATE TABLE
  email_verifications (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    user_id BIGINT REFERENCES users (id) ON DELETE SET NULL,
    verification_token VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL,
    attempt_count SMALLINT NOT NULL DEFAULT 0,
    max_attempts SMALLINT NOT NULL DEFAULT 5,
    verified_at TIMESTAMP,
    expires_at BIGINT NOT NULL,
    is_revoked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

CREATE INDEX "idx_email_verifications_user" ON email_verifications (user_id, expires_at, is_revoked)