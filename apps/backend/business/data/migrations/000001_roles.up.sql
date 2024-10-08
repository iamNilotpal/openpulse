-- Enum for role types
CREATE TYPE IF NOT EXISTS app_role_type AS ENUM (
  'team_admin',
  'team_billing_admin',
  'team_lead',
  'team_responder',
  'team_member',
);

-- Enum for action types
CREATE TYPE IF NOT EXISTS permission_action_type AS ENUM ('view', 'create', 'update', 'delete', 'manage');

-- Enum for resource types
CREATE TYPE NOT EXISTS resource_type AS ENUM (
  'teams',
  'team_members',
  'billings',
  'global_api_tokens',
  'team_api_tokens',
  'monitors',
  'heartbeats',
  'integrations',
  'incidents',
  'invitations',
  'status_pages',
  'escalation_policy',
  'on_call_calendars',
  'sources',
  'dashboards'
);

CREATE TABLE
  IF NOT EXISTS roles (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(80) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    role app_role_type NOT NULL,
    is_system_role BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX (role)
  );

CREATE TABLE
  IF NOT EXISTS resources (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    resource resource_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX (resource)
  );

CREATE TABLE
  IF NOT EXISTS roles_resources (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    role_id SMALLINT NOT NULL REFERENCES roles (id) ON DELETE CASCADE,
    resource_id SMALLINT NOT NULL REFERENCES resources (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (role_id, resource_id),
    INDEX (role_id)
  );

CREATE TABLE
  IF NOT EXISTS permissions (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(80) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    action permission_action_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX (action)
  );

CREATE TABLE
  IF NOT EXISTS resource_permissions (
    id SMALLSERIAL PRIMARY KEY NOT NULL,
    resource_id SMALLINT NOT NULL REFERENCES resources (id) ON DELETE CASCADE,
    permission_id SMALLINT NOT NULL REFERENCES permissions (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (resource_id, permission_id),
    INDEX (resource_id)
  );

--
-- ROLES:
--   1. Admin: Can change billing, dashboards, global API tokens, heartbeats, incidents, integrations,
--   monitors, on-call calendars, policies, severities, sources, status pages, team
--   members, teams, and the organization.
--
--   2. Billing Admin: Can access the organization. Can change billing.
--
--   3. Team Lead: Can access billing, global API tokens, teams, and the organization. Can change
--   dashboards, heartbeats, incidents, integrations, monitors, on-call calendars, policies,
--   severities, sources, status pages, and team members.
--
--   4. Responder: Can access billing, team members, teams, and the organization. Can change
--   dashboards, heartbeats, incidents, integrations, monitors, on-call calendars, policies,
--   severities, sources, and status pages.
--
--   5. Member: Can access billing, team members, teams, and the organization. Can change dashboards,
--   and sources.
--
-- RESOURCES:
--   Teams, Billing, Team Members, Global API Tokens, Monitors, Heartbeats, Integrations, On
--   call calendars, Incidents, Escalation and Security Levels, Status Page, Dashboards.
--