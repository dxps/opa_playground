package domain

import (
	"errors"
	"time"

	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

// Subject represents a human individual or a service (machine) account
// that is registered within this system and used for authentication.
type Subject struct {
	IID       int64     `json:"-"`
	EID       string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Active    bool      `json:"active"`
	Version   int       `json:"-"`
}

type password struct {
	plaintext *string
	Hash      []byte
}

// Set does the password plain and hash versions set.
func (p *password) Set(plaintext string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return err
	}
	p.plaintext = &plaintext
	p.Hash = hash
	return nil
}

func (p *password) Matches(plaintext string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plaintext))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil

		default:
			return false, err
		}
	}
	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRegexp), "email", "must be a valid address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

func ValidateSubject(v *validator.Validator, subj *Subject) {

	v.Check(subj.Name != "", "name", "must be provided")
	v.Check(len(subj.Name) <= 500, "name", "must not be more than 500 bytes long")

	ValidateEmail(v, subj.Email)

	if subj.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *subj.Password.plaintext)
	}

	// Normally, that shouldn't be the case, but better to have it covered
	// by raising a panic since it reflects an internal bug, not a validation error.
	if subj.Password.Hash == nil {
		panic("missing password hash of the user")
	}
}
