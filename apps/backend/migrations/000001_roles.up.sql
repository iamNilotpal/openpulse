CREATE TABLE
  IF NOT EXISTS roles (
    id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    name VARCHAR(80) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    role SMALLINT UNIQUE NOT NULL CHECK (role > 0),
    is_system_role BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

CREATE INDEX "idx_roles_role" ON roles (role);

CREATE INDEX "idx_roles_role_name" ON roles (name);

CREATE TABLE
  IF NOT EXISTS resources (
    id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    resource SMALLINT UNIQUE NOT NULL CHECK (resource > 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

CREATE INDEX "idx_resources_resource" ON resources (resource);

CREATE TABLE
  IF NOT EXISTS roles_resources (
    id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    role_id SMALLINT NOT NULL REFERENCES roles (id) ON DELETE NO ACTION,
    resource_id SMALLINT NOT NULL REFERENCES resources (id) ON DELETE NO ACTION,
    UNIQUE (role_id, resource_id)
  );

CREATE INDEX "idx_roles_resources_role_id" ON roles_resources (role_id);

CREATE TABLE
  IF NOT EXISTS permissions (
    id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    name VARCHAR(80) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    action SMALLINT UNIQUE NOT NULL CHECK (action > 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS resource_permissions (
    id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    resource_id SMALLINT NOT NULL REFERENCES resources (id) ON DELETE NO ACTION,
    permission_id SMALLINT NOT NULL REFERENCES permissions (id) ON DELETE NO ACTION,
    UNIQUE (resource_id, permission_id)
  );

CREATE INDEX "idx_resource_permissions_resource_id" ON resource_permissions (resource_id);