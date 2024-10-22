package auth_handlers

import "github.com/iamNilotpal/openpulse/foundation/validate"

type SignUpInput struct {
	FirstName string `json:"firstName" validate:"required,min=1,max=255"`
	LastName  string `json:"lastName" validate:"required,min=1,max=255"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=50"`
}

func (v SignUpInput) Validate() error {
	return validate.Check(v)
}

type RegisterUserResponse struct {
	URL    string            `json:"url"` // Temporary
	State  RegistrationState `json:"state,omitempty"`
	UserId int               `json:"userId,omitempty"`
}

type RegistrationState struct {
	EmailSent     string `json:"emailSentState"`
	CreateAccount string `json:"createAccountState"`
}

type SignInInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (v SignInInput) Validate() error {
	return validate.Check(v)
}

type SignInResponse struct {
	UserId       int    `json:"userId"`
	SessionId    int    `json:"sessionId"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type IncompleteOnboardingResponse struct {
	UserId           int    `json:"userId"`
	AccessToken      string `json:"accessToken"`
	OnboardingStatus string `json:"onboardingStatus"`
}

type NewOauthUserInput struct {
	FirstName string `json:"firstName" validate:"required,min=1,max=255"`
	LastName  string `json:"lastName" validate:"required,min=1,max=255"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"phone"`
	AvatarURL string `json:"avatarURL" validate:"url"`
}

type NewOAuthAccountInput struct {
	ExternalId string            `json:"externalId" validate:"required,min=1"`
	Scope      string            `json:"scope" validate:"required,min=1"`
	Metadata   string            `json:"metadata"`
	User       NewOauthUserInput `json:"user" validate:"required"`
}

func (v NewOAuthAccountInput) Validate() error {
	return validate.Check(v)
}
