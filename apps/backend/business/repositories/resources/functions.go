package resources

import (
	"time"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	resources_store "github.com/iamNilotpal/openpulse/business/repositories/resources/store/postgres"
)

func FromDBResource(cmd resources_store.Resource) Resource {
	createdAt, _ := time.Parse(time.UnixDate, cmd.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, cmd.UpdatedAt)

	return Resource{
		Id:          cmd.Id,
		Name:        cmd.Name,
		Description: cmd.Description,
		Resource:    ParseAppResourceInt(cmd.Resource),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func ToNewDBResource(cmd NewResource) resources_store.NewResource {
	return resources_store.NewResource{
		Name:        cmd.Name,
		Description: cmd.Description,
		Resource:    ParseAppResource(cmd.Resource),
	}
}

func FromDBResourceAccessDetails(cmd resources_store.ResourceAccessConfig) ResourceAccessConfig {
	return ResourceAccessConfig{
		Id:       cmd.Id,
		Resource: ParseAppResourceInt(cmd.Resource),
	}
}

func FromDBResourceWithPermission(cmd resources_store.ResourceWithPermission) ResourceWithPermission {
	return ResourceWithPermission{
		Resource:   FromDBResourceAccessDetails(cmd.Resource),
		Permission: permissions.FromDBPermissionAccessDetails(cmd.Permission),
	}
}
