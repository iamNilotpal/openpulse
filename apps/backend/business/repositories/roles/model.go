package roles

import (
	"time"

	modified_by "github.com/iamNilotpal/openpulse/business/data/modified-by"
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/postgres"
)

type Role struct {
	Id           int
	IsSystemRole bool
	Name         string
	Description  string
	UpdatedBy    modified_by.ModifiedBy
	CreatedBy    modified_by.ModifiedBy
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RoleAccessConfig struct {
	Id   int
	Role AppRole
}

type RoleAccessControl struct {
	Role       RoleAccessConfig
	Resource   resources.ResourceAccessConfig
	Permission permissions.PermissionAccessConfig
}

type NewRole struct {
	CreatorId    int
	IsSystemRole bool
	Name         string
	Description  string
}

type UpdateRole struct {
	Name        string
	Description string
}

func ConstructRole(name, description string, creatorId int, isSystemRole bool) NewRole {
	return NewRole{
		Name:         name,
		CreatorId:    creatorId,
		Description:  description,
		IsSystemRole: isSystemRole,
	}
}

func ToNewDBRole(r NewRole) roles_store.NewRole {
	return roles_store.NewRole{
		Name:         r.Name,
		CreatorId:    r.CreatorId,
		Description:  r.Description,
		IsSystemRole: r.IsSystemRole,
	}
}

func FromDBRole(r roles_store.Role) Role {
	createdAt, _ := time.Parse(time.UnixDate, r.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, r.UpdatedAt)

	return Role{
		Id:           r.Id,
		Name:         r.Name,
		Description:  r.Description,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		IsSystemRole: r.IsSystemRole,
		UpdatedBy: modified_by.New(
			r.UpdatedBy.Id, r.UpdatedBy.Email, r.UpdatedBy.FirstName, r.UpdatedBy.LastName,
		),
		CreatedBy: modified_by.New(
			r.CreatedBy.Id, r.CreatedBy.Email, r.CreatedBy.FirstName, r.CreatedBy.LastName,
		),
	}
}

func FromDBRoleAccessConfig(r roles_store.RoleAccessConfig) RoleAccessConfig {
	return RoleAccessConfig{
		Id:   r.Id,
		Role: ToAppRole(r.Role),
	}
}

func FromDBRoleAccessControl(r roles_store.RoleAccessControl) RoleAccessControl {
	return RoleAccessControl{
		Role:       FromDBRoleAccessConfig(r.Role),
		Resource:   resources.FromDBResourceAccessDetails(r.Resource),
		Permission: permissions.FromDBPermissionAccessDetails(r.Permission),
	}
}
