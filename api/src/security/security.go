package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(passwordWithHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(passwordString))
}
