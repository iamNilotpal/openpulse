package emails_store

type EmailVerificationInput struct {
	UserId            int
	ExpiresAt         int
	Email             string
	VerificationToken string
}
