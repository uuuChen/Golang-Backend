package users

import (
	"net/mail"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPasswordFormat(password string) bool {
	if len(password) < 6 || len(password) > 16 {
		return false
	}

	var uppercase, lowercase, number, special bool
	specialRegex := regexp.MustCompile(`[()\[\]{}<>\+\-\*/?,.:;"'_\|~` + "`" + `!@#\$%\^&=]`)
	for _, c := range password {
		switch {
		case c >= 'A' && c <= 'Z':
			uppercase = true
		case c >= 'a' && c <= 'z':
			lowercase = true
		case c >= '0' && c <= '9':
			number = true
		case specialRegex.MatchString(string(c)):
			special = true
		}
	}
	return uppercase && lowercase && number && special
}

// Ref: https://medium.com/@rnp0728/secure-password-hashing-in-go-a-comprehensive-guide-5500e19e7c1f
func HashPlainPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// Ref: https://medium.com/@rnp0728/secure-password-hashing-in-go-a-comprehensive-guide-5500e19e7c1f
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
