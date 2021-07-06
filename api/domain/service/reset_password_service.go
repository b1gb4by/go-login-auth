package service

import (
	"api/util"
	"fmt"
	"math/rand"
	"net/smtp"
)

func CreateRandomString() string {
	const n = 36
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

func SendToSMTP(to string, token string) error {
	const from = "password-reset@example.com"
	const mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	sendFrom := fmt.Sprintf("From: %s\n", from)
	subject := fmt.Sprintf("Subject; %s\n", "Password Reset")
	url := "http://localhost:8080/reset/" + token
	message := fmt.Sprintf("Click <a href=\"%s\">here</a> to reset password!", url)

	if err := smtp.SendMail(
		"smtp:1025",
		nil,
		from,
		[]string{to},
		[]byte(sendFrom+subject+mime+message),
	); err != nil {
		return util.Errorf(util.ErrorCode10008, "", "%w", err)
	}
	return nil
}
