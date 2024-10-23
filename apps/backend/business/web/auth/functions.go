package auth

import (
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
)

/* ================ User Functions ================ */
func NewUserRoleConfig(role users.Role) UserRoleConfig {
	return UserRoleConfig{Id: role.Id, Role: role.Role}
}

func NewUserPermissionConfig(p users.Permission) UserPermissionConfig {
	return UserPermissionConfig{Id: p.Id, Action: p.Action, Enabled: p.Enabled}
}

func NewUserResourceConfig(r users.Resource) UserResourceConfig {
	return UserResourceConfig{Id: r.Id, Resource: r.Resource}
}

/* ================ App Functions ================ */
func NewRoleConfig(role roles.RoleAccessConfig) RoleConfig {
	return RoleConfig{Id: role.Id, Role: role.Role}
}

func NewPermissionConfig(p permissions.PermissionAccessConfig) PermissionConfig {
	return PermissionConfig{Id: p.Id, Action: p.Action}
}

func NewResourceConfig(r resources.ResourceAccessConfig) ResourceConfig {
	return ResourceConfig{Id: r.Id, Resource: r.Resource}
}

func BuildAuthorizationMaps(
	r []roles.Role, res []resources.Resource, perms []permissions.Permission,
) (RoleMappings, ResourceMappings, PermissionMappings) {
	roleIdMap := make(RoleIDMap)
	roleNameMap := make(RoleNameMap)

	for _, role := range r {
		roleIdMap[role.Id] = NewRoleConfig(roles.RoleAccessConfig{Id: role.Id, Role: role.Role})
		roleNameMap[role.Role] = NewRoleConfig(roles.RoleAccessConfig{Id: role.Id, Role: role.Role})
	}

	resourceTypeMap := make(ResourceTypeMap)
	resourceTypeIdMap := make(ResourceTypeIdMap)

	for _, res := range res {
		resourceTypeMap[res.Resource] = NewResourceConfig(
			resources.ResourceAccessConfig{Id: res.Id, Resource: res.Resource},
		)
		resourceTypeIdMap[res.Id] = NewResourceConfig(
			resources.ResourceAccessConfig{Id: res.Id, Resource: res.Resource},
		)
	}

	permActionMap := make(PermissionActionMap)
	permActionIdMap := make(PermissionActionIdMap)

	for _, p := range perms {
		permActionMap[p.Action] = NewPermissionConfig(
			permissions.PermissionAccessConfig{Id: p.Id, Action: p.Action},
		)
		permActionIdMap[p.Id] = NewPermissionConfig(
			permissions.PermissionAccessConfig{Id: p.Id, Action: p.Action},
		)
	}

	return RoleMappings{ByID: roleIdMap, ByName: roleNameMap},
		ResourceMappings{ByName: resourceTypeMap, ByID: resourceTypeIdMap},
		PermissionMappings{ByAction: permActionMap, ByID: permActionIdMap}
}

func BuildAccessControlMaps(roleAccessControls []roles.RoleAccessControl) (
	ResourceToPermissionsMap, RoleNameToAccessControlMap,
) {
	resourceToPermissionsMap := make(ResourceToPermissionsMap)
	roleNameToAccessControlMap := make(RoleNameToAccessControlMap)

	for _, rac := range roleAccessControls {
		storedResToPermsMap, ok := roleNameToAccessControlMap[rac.Role.Role]
		if !ok {
			resPermsMap := make(ResourceToPermissionsMap)
			resPermsMap[rac.Resource.Resource] = ResourcePermConfig{
				Resource:    NewResourceConfig(rac.Resource),
				Permissions: []PermissionConfig{NewPermissionConfig(rac.Permission)},
			}

			roleNameToAccessControlMap[rac.Role.Role] = resPermsMap
			resourceToPermissionsMap = resPermsMap
			continue
		}

		resourcePermissions, ok := storedResToPermsMap[rac.Resource.Resource]
		if !ok {
			storedResToPermsMap[rac.Resource.Resource] = ResourcePermConfig{
				Resource:    NewResourceConfig(rac.Resource),
				Permissions: []PermissionConfig{NewPermissionConfig(rac.Permission)},
			}
			roleNameToAccessControlMap[rac.Role.Role] = storedResToPermsMap
			resourceToPermissionsMap = storedResToPermsMap
			continue
		}

		resourcePermissions.Permissions = append(
			resourcePermissions.Permissions, NewPermissionConfig(rac.Permission),
		)
		storedResToPermsMap[rac.Resource.Resource] = resourcePermissions
		roleNameToAccessControlMap[rac.Role.Role] = storedResToPermsMap
		resourceToPermissionsMap = storedResToPermsMap
	}

	return resourceToPermissionsMap, roleNameToAccessControlMap
}
