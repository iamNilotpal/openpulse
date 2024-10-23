package roles_handler

import "github.com/iamNilotpal/openpulse/foundation/validate"

type NewRoleInput struct {
	Role        string `validate:"required,min=1,max=50" json:"role"`
	Name        string `validate:"required,min=1,max=80" json:"name"`
	Description string `validate:"required,min=10,max=255" json:"description"`
}

func (v NewRoleInput) Validate() error {
	return validate.Check(v)
}

type NewRoleResponse struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
	Name string `json:"name"`
}

type AppRole struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
