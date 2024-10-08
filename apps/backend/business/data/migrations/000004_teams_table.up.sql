CREATE TABLE
  IF NOT EXISTS organizations (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    logo_url TEXT,
    total_employees TEXT,
    admin_id BIGINT NOT NULL REFERENCES users (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX (admin_id)
  );

CREATE TABLE
  IF NOT EXISTS teams (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    total_members SMALLINT NOT NULL,
    logo_url TEXT,
    invitation_code VARCHAR(100) UNIQUE NOT NULL,
    creator_id BIGINT NOT NULL REFERENCES users (id),
    org_id BIGINT NOT NULL REFERENCES organizations (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX (creator_id),
    INDEX (org_id)
  );

CREATE TABLE
  IF NOT EXISTS team_users (
    id BIGSERIAL PRIMARY KEY NOT NULl,
    team_id BIGINT NOT NULL REFERENCES teams (id),
    user_id BIGINT NOT NULL REFERENCES users (id),
    role_id SMALLINT NOT NULL REFERENCES roles (id),
    UNIQUE (team_id, user_id, role_id),
    INDEX (team_id),
    INDEX (user_id)
  );

ALTER TABLE users
ADD COLUMN team_id BIGINT REFERENCES teams (id);