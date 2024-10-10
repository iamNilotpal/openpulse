package auth_handlers

import "github.com/iamNilotpal/openpulse/business/sys/validate"

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required,min=1,max=255"`
	LastName  string `json:"lastName" validate:"required,min=1,max=255"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=5,max=50"`
}

func (rup RegisterUserPayload) Validate() error {
	return validate.Check(rup)
}

type RegisterUserResponse struct {
	UserId int `json:"userId"`
}
