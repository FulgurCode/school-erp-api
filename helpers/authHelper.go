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

// Hashing password using bcrypt
func HashPassword(password string) string {
	// Generating hashed password
	var hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
