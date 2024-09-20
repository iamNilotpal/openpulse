CREATE TABLE
  IF NOT EXISTS teams (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    total_members SMALLINT NOT NULL,
    admin_id BIGINT NOT NULL REFERENCES users (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS team_users (
    team_id BIGINT NOT NULL REFERENCES teams (id),
    user_id BIGINT NOT NULL REFERENCES users (id),
    role_id BIGINT NOT NULL REFERENCES roles (id),
    UNIQUE (team_id, user_id, role_id),
    INDEX (team_id, user_id, role_id)
  );