package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plaintext string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)

	return string(hp), err
}
