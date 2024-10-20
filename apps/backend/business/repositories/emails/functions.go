package emails

import emails_store "github.com/iamNilotpal/openpulse/business/repositories/emails/store/postgres"

func NewDBEmailVerificationDetails(v EmailVerificationDetails) emails_store.EmailVerificationInput {
	return emails_store.EmailVerificationInput{
		Email:             v.Email,
		UserId:            v.UserId,
		VerificationToken: v.VerificationToken,
		ExpiresAt:         int(v.ExpiresAt.UnixNano()),
	}
}
