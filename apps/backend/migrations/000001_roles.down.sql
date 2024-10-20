-- Drop the resource_permissions table
DROP TABLE IF EXISTS resource_permissions;

-- Drop the permissions table
DROP TABLE IF EXISTS permissions;

-- Drop the roles_resources table
DROP TABLE IF EXISTS roles_resources;

-- Drop the resources table
DROP TABLE IF EXISTS resources;

-- Drop the roles table
DROP TABLE IF EXISTS roles;

-- Drop the enums (in reverse order)
DROP TYPE IF EXISTS resource_type;

DROP TYPE IF EXISTS permission_action_type;

DROP TYPE IF EXISTS app_role_type;