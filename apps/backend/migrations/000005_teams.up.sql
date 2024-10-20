CREATE TABLE
  IF NOT EXISTS organizations (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    logo_url TEXT,
    total_employees TEXT NOT NULL,
    admin_id BIGINT NOT NULL REFERENCES users (id) ON DELETE NO ACTION,
    deleted_at BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (admin_id)
  );

CREATE INDEX idx_organizations_deleted_at ON organizations (deleted_at);

CREATE TABLE
  IF NOT EXISTS teams (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    total_members SMALLINT NOT NULL,
    logo_url TEXT,
    invitation_code VARCHAR(100) UNIQUE NOT NULL,
    creator_id BIGINT NOT NULL REFERENCES users (id) ON DELETE NO ACTION,
    org_id BIGINT NOT NULL REFERENCES organizations (id) ON DELETE NO ACTION,
    deleted_at BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

ALTER TABLE users
ADD COLUMN team_id BIGINT REFERENCES teams (id) ON DELETE SET NULL;

CREATE INDEX "idx_teams_creator_id" ON teams (creator_id);

CREATE INDEX "idx_teams_org_id" ON teams (org_id);

CREATE INDEX idx_teams_deleted_at ON teams (deleted_at);

CREATE TABLE
  IF NOT EXISTS team_users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULl,
    team_id BIGINT NOT NULL REFERENCES teams (id) ON DELETE NO ACTION,
    user_id BIGINT NOT NULL REFERENCES users (id) ON DELETE NO ACTION,
    role_id SMALLINT NOT NULL REFERENCES roles (id) ON DELETE NO ACTION,
    resource_id SMALLINT NOT NULL REFERENCES resources (id) ON DELETE NO ACTION,
    permission_id SMALLINT NOT NULL REFERENCES permissions (id) ON DELETE NO ACTION,
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    UNIQUE (
      team_id,
      user_id,
      role_id,
      resource_id,
      permission_id
    )
  );

CREATE INDEX "idx_team_users_team_id" ON team_users (team_id);

CREATE INDEX "idx_team_users_user_id" ON team_users (user_id);

CREATE TABLE
  IF NOT EXISTS team_invitations (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    email VARCHAR(255) NOT NULL,
    role_id BIGINT NOT NULL REFERENCES roles (id) ON DELETE NO ACTION,
    team_id BIGINT NOT NULL REFERENCES teams (id) ON DELETE NO ACTION,
    org_id BIGINT NOT NULL REFERENCES organizations (id) ON DELETE NO ACTION,
    invited_by BIGINT NOT NULL REFERENCES users (id) ON DELETE NO ACTION,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_in BIGINT NOT NULL,
    status SMALLINT DEFAULT 1 NOT NULL CHECK (status > 0), -- pending, accepted, expired
    invited_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    accepted_at TIMESTAMP
  );

CREATE INDEX "idx_team_invitations_expires_in" ON team_invitations (expires_in);