package auth_handlers

import "github.com/iamNilotpal/openpulse/business/sys/validate"

type SignUpPayload struct {
	FirstName string `json:"firstName" validate:"required,min=1,max=255"`
	LastName  string `json:"lastName" validate:"required,min=1,max=255"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=50"`
}

func (sup SignUpPayload) Validate() error {
	return validate.Check(sup)
}

type RegisterUserResponse struct {
	UserId int `json:"userId"`
}

type SignInPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (sip SignInPayload) Validate() error {
	return validate.Check(sip)
}
