package auth_handlers

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=5"`
}

type RegisterUserResponse struct {
	UserId int `json:"userId"`
}

type VerifyEmailPayload struct {
	Token string `json:"token"`
}
