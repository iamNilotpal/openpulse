package roles_handler

import "github.com/iamNilotpal/openpulse/business/sys/validate"

type NewAppRole struct {
	Name        string `validate:"required,min=3" json:"name"`
	Description string `validate:"required,min=10" json:"description"`
}

type AppRole struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func (app NewAppRole) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}

	return nil
}
