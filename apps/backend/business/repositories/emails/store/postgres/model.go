package emails_store

type EmailVerificationInput struct {
	UserId            int
	MaxAttempts       int
	ExpiresAt         int
	Email             string
	VerificationToken string
}
