package helpers

import "golang.org/x/crypto/bcrypt"

// Comparing password and creating response
func ComparePassword(hashedPassword string, password string) bool {
	var err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
