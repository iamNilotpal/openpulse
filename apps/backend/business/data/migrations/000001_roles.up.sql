CREATE TABLE
  IF NOT EXISTS roles (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(80) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS permissions (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(80) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    action VARCHAR(20) NOT NULL CHECK (action IN ('create', 'update', 'delete', 'view')),
    resource VARCHAR(30) NOT NULL CHECK (
      resource IN (
        'teams',
        'monitor',
        'members',
        'billing',
        'incidents',
        'heartbeats',
        'invitations',
        'status_pages',
        'escalation_policy',
        'on_call_escalations',
      )
    ),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS roles_permissions (
    role_id SMALLINT NOT NULL REFERENCES roles (id),
    permission_id SMALLINT NOT NULL REFERENCES permissions (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (role_id, permission_id),
    INDEX (role_id) INCLUDE (permission_id)
  );