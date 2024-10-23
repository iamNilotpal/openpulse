package resources_handlers

import "github.com/iamNilotpal/openpulse/foundation/validate"

type NewResourceInput struct {
	Resource    string `json:"resource" validate:"required,min=1"`
	Name        string `json:"name" validate:"required,min=1,max=80"`
	Description string `json:"description" validate:"required,min=1,max=300"`
}

func (v NewResourceInput) Validate() error {
	return validate.Check(v)
}

type NewResourceResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Resource string `json:"resource"`
}
