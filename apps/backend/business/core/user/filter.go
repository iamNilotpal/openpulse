package user

import (
	"net/mail"

	"github.com/iamNilotpal/openpulse/business/sys/validate"
)

type QueryFilter struct {
	ID        int          `validate:"omitempty"`
	Email     mail.Address `validate:"omitempty"`
	FirstName string       `validate:"omitempty,min=3"`
	LastName  string       `validate:"omitempty,min=3"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return err
	}
	return nil
}

func (qf *QueryFilter) WithUserID(userID int) {
	qf.ID = userID
}

func (qf *QueryFilter) WithFirstName(firstName string) {
	qf.FirstName = firstName
}

func (qf *QueryFilter) WithLastName(lastName string) {
	qf.LastName = lastName
}

func (qf *QueryFilter) WithEmail(email mail.Address) {
	qf.Email = email
}
