CREATE TABLE
  IF NOT EXISTS invitations (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    email VARCHAR(255) NOT NULL,
    role_id BIGINT NOT NULL REFERENCES roles (id),
    team_id BIGINT NOT NULL REFERENCES teams (id),
    invited_by BIGINT NOT NULL REFERENCES users (id),
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_in BIGINT NOT NULL,
    invitation_status VARCHAR(15) CHECK (
      invitation_status IN ('pending', 'accepted', 'expired')
    ) DEFAULT 'pending' NOT NULL,
    invited_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    accepted_at TIMESTAMP,
    INDEX (expires_in) INCLUDE (id)
  );