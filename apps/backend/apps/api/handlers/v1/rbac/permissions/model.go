package permissions_handlers

import "github.com/iamNilotpal/openpulse/foundation/validate"

type NewPermissionInput struct {
	Name        string `json:"name" validate:"required,min=1,max=80"`
	Description string `json:"description" validate:"required,min=1,max=300"`
	Action      string `json:"action" validate:"required,oneof=view create update delete manage"`
}

func (v NewPermissionInput) Validate() error {
	return validate.Check(v)
}

type NewPermissionResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Action string `json:"action"`
}
