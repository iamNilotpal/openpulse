package auth_handlers

import "github.com/iamNilotpal/openpulse/business/sys/validate"

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=5"`
}

func (rup RegisterUserPayload) Validate() error {
	return validate.Check(rup)
}

type RegisterUserResponse struct {
	UserId int `json:"userId"`
}

type VerifyEmailPayload struct {
	Token string `json:"token"`
}
