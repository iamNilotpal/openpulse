package resources

import (
	"time"

	modified_by "github.com/iamNilotpal/openpulse/business/data/modified-by"
	resources_store "github.com/iamNilotpal/openpulse/business/repositories/resources/store/postgres"
)

type Resource struct {
	Id        int
	Name      string
	Resource  string
	CreatedBy modified_by.ModifiedBy
	UpdatedBy modified_by.ModifiedBy
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewResource struct {
	CreatorId int
	Name      string
	Resource  ResourceType
}

func ToNewDBResource(r NewResource) resources_store.NewResource {
	return resources_store.NewResource{
		Name:      r.Name,
		CreatorId: r.CreatorId,
		Resource:  FromResourceType(r.Resource),
	}
}

func FromDBResource(r resources_store.Resource) Resource {
	createdAt, _ := time.Parse(time.UnixDate, r.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, r.UpdatedAt)

	return Resource{
		Id:        r.Id,
		Name:      r.Name,
		Resource:  r.Resource,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		CreatedBy: modified_by.New(
			r.CreatedBy.Id, r.CreatedBy.Email, r.CreatedBy.FirstName, r.CreatedBy.LastName,
		),
		UpdatedBy: modified_by.New(
			r.UpdatedBy.Id, r.UpdatedBy.Email, r.UpdatedBy.FirstName, r.UpdatedBy.LastName,
		),
	}
}
