package roles_handler

import "github.com/iamNilotpal/openpulse/business/sys/validate"

type NewRole struct {
	Role        string `validate:"required,min=1,max=50" json:"role"`
	Name        string `validate:"required,min=1,max=255" json:"name"`
	Description string `validate:"required,min=10,max=255" json:"description"`
}

func (na NewRole) Validate() error {
	return validate.Check(na)
}

type AppRole struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
