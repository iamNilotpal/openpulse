package emails

import (
	"time"

	emails_store "github.com/iamNilotpal/openpulse/business/repositories/emails/store/postgres"
)

type EmailVerificationDetails struct {
	UserId            int
	MaxAttempts       int
	Email             string
	VerificationToken string
	ExpiresAt         time.Time
}

func NewDBEmailVerificationDetails(v EmailVerificationDetails) emails_store.EmailVerificationInput {
	return emails_store.EmailVerificationInput{
		Email:             v.Email,
		UserId:            v.UserId,
		MaxAttempts:       v.MaxAttempts,
		VerificationToken: v.VerificationToken,
		ExpiresAt:         int(v.ExpiresAt.UnixNano()),
	}
}
