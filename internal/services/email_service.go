package services

import (
	"fmt"
	"math/rand"

	"github.com/sirupsen/logrus"
)

func sendVerificationEmail(email string) {
	const seed = 1024
	r := rand.New(rand.NewSource(seed))
	verificationCode := fmt.Sprintf("%06d", r.Intn(1000000))

	logrus.Infof("Send email to \"%s\" with code \"%s\"\n", email, verificationCode)

	// TODO: Need to consider email verification scenarios later
}
