DROP TABLE IF EXISTS teams;

DROP TABLE IF EXISTS team_users;

SELECT
  r.id AS role_id,
  r.name AS role_name,
  r.description AS role_description,
  r.is_system_role AS is_system_role,
  r.created_at AS role_created_at,
  r.updated_at AS role_updated_at,
  rcb.first_name AS role_author_first_name,
  rcb.last_name AS role_author_last_name,
  rub.first_name AS role_updater_last_name,
  rub.last_name AS role_updated_last_name,
  --
  p.id AS permission_id,
  p.name AS permission_name,
  p.description AS permission_description,
  p.action AS permission_action,
  p.created_at AS permission_created_at,
  p.updated_at AS permission_updated_at,
  pcb.first_name AS permission_author_first_name,
  pcb.last_name AS permission_author_last_name,
  pub.first_name AS permission_updater_last_name,
  pub.last_name AS permission_updated_last_name,
  --
  rp.created_at AS role_permission_created_at,
  rp.updated_at AS role_permission_updated_at,
  rpcb.first_name AS role_permission_author_first_name,
  rpcb.last_name AS role_permission_author_last_name,
  rpub.first_name AS role_permission_updater_last_name,
  rpub.last_name AS role_permission_updated_last_name
FROM
  roles r
  JOIN roles_permissions rp ON rp.role_id = r.id
  JOIN permissions p ON p.id = rp.permission_id
  -- Join users table for roles
  LEFT JOIN users rcb ON rcb.id = r.created_by
  LEFT JOIN users rub ON rub.id = r.updated_by
  -- Join users table for permissions
  LEFT JOIN users pcb ON pcb.id = p.created_by
  LEFT JOIN users pub ON pub.id = p.updated_by
  -- Join users table for roles_permissions
  LEFT JOIN users rpcb ON rpcb.id = rp.created_by
  LEFT JOIN users rpub ON rpub.id = rp.updated_by;