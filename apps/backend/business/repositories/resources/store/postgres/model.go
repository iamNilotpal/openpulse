package resources_store

import modified_by "github.com/iamNilotpal/openpulse/business/data/modified-by"

type Resource struct {
	Id        int
	Name      string
	Resource  string
	CreatedBy modified_by.ModifiedBy
	UpdatedBy modified_by.ModifiedBy
	CreatedAt string
	UpdatedAt string
}

type NewResource struct {
	CreatorId int
	Name      string
	Resource  string
}
