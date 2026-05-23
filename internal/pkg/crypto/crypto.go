package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt.
func HashPassword(password string, cost ...int) (string, error) {
	c := bcrypt.DefaultCost
	if len(cost) > 0 {
		c = cost[0]
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), c)
	return string(bcryptPassword), err
}

// CheckPasswordHash verifies a password against a bcrypt hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
