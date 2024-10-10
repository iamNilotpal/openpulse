package resources

import (
	"time"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	resources_store "github.com/iamNilotpal/openpulse/business/repositories/resources/store/postgres"
)

type Resource struct {
	Id          int
	Name        string
	Description string
	Resource    AppResource
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewResource struct {
	Name        string
	Description string
	Resource    AppResource
}

type ResourceAccessConfig struct {
	Id       int
	Resource AppResource
}

type ResourceWithPermission struct {
	Resource   ResourceAccessConfig
	Permission permissions.PermissionAccessConfig
}

func FromDBResource(r resources_store.Resource) Resource {
	createdAt, _ := time.Parse(time.UnixDate, r.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, r.UpdatedAt)

	return Resource{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
		Resource:    ParseAppResourceInt(r.Resource),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func ToNewDBResource(r NewResource) resources_store.NewResource {
	return resources_store.NewResource{
		Name:        r.Name,
		Description: r.Description,
		Resource:    ParseAppResource(r.Resource),
	}
}

func FromDBResourceAccessDetails(r resources_store.ResourceAccessConfig) ResourceAccessConfig {
	return ResourceAccessConfig{
		Id:       r.Id,
		Resource: ParseAppResourceInt(r.Resource),
	}
}

func FromDBResourceWithPermission(r resources_store.ResourceWithPermission) ResourceWithPermission {
	return ResourceWithPermission{
		Resource:   FromDBResourceAccessDetails(r.Resource),
		Permission: permissions.FromDBPermissionAccessDetails(r.Permission),
	}
}
