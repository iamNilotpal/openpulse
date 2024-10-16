package auth_handlers

import "github.com/iamNilotpal/openpulse/business/sys/validate"

type SignUpInput struct {
	FirstName string `json:"firstName" validate:"required,min=1,max=255"`
	LastName  string `json:"lastName" validate:"required,min=1,max=255"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=50"`
}

func (sup SignUpInput) Validate() error {
	return validate.Check(sup)
}

type RegisterUserResponse struct {
	State  RegistrationState `json:"state,omitempty"`
	UserId int               `json:"userId,omitempty"`
}

type SignInInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (sip SignInInput) Validate() error {
	return validate.Check(sip)
}

type RegistrationState struct {
	EmailSentState    string `json:"emailSentState"`
	RegistrationState string `json:"registrationState"`
}
