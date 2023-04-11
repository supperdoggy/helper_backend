package utils

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/supperdoggy/helper/pkg/models"
	"github.com/supperdoggy/helper/pkg/models/dbmodels"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

const minEntropyBits = 50

// ValidateUserEmailAndPassword validates the email and password, returns error if invalid
func ValidateUserEmailAndPassword(email, password string) error {
	if err := validation.Validate(&email, validation.Required, is.Email); err != nil {
		return errors.New("invalid email")
	}

	// validate password
	err := passwordvalidator.Validate(password, minEntropyBits)
	if err != nil {
		return errors.New("password is too simple")
	}
	return nil
}

// MapDBUserToResponseUser takes an input of type dbmodels.User and
// returns a pointer to models.User struct.
// The new instance has ID, Email, CreatedAt, and EditedAt fields set to corresponding values of the input
func MapDBUserToResponseUser(u dbmodels.User) *models.User {
	return &models.User{
		ID:       u.ID,
		Email:    u.Email,
		FullName: u.FullName,

		CreatedAt: u.CreatedAt,
		EditedAt:  u.EditedAt,
	}
}
