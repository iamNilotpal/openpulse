package emails

import (
	"time"
)

type EmailVerificationDetails struct {
	UserId            int
	MaxAttempts       int
	Email             string
	VerificationToken string
	ExpiresAt         time.Time
}
