CREATE TABLE
  IF NOT EXISTS roles (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(80) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    role SMALLINT UNIQUE NOT NULL,
    is_system_role BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS resources (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    resource SMALLINT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS roles_resources (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    role_id SMALLINT NOT NULL REFERENCES roles (id) ON DELETE CASCADE ON UPDATE CASCADE,
    resource_id SMALLINT NOT NULL REFERENCES resources (id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (role_id, resource_id)
  );

CREATE TABLE
  IF NOT EXISTS permissions (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(80) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    action SMALLINT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS resource_permissions (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    resource_id SMALLINT NOT NULL REFERENCES resources (id) ON DELETE CASCADE,
    permission_id SMALLINT NOT NULL REFERENCES permissions (id) ON DELETE CASCADE,
    UNIQUE (resource_id, permission_id)
  );

CREATE INDEX "idx_roles_role" ON roles (role);

CREATE INDEX "idx_resources_resource" ON resources (resource);

CREATE INDEX "idx_roles_resources_role_id" ON roles_resources (role_id);

CREATE INDEX "idx_resource_permissions_resource_id" ON resource_permissions (resource_id);