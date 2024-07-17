package redis

import (
	"AuthService/config"
	"context"
	"fmt"
	"net/smtp"
)

func SendVerificationToEmail(email, code string) error {
	rdb, err := ConnectRedis()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}
	from := "azizbekqobulov05@gmail.com"
	password := config.Load().GMAIL_CODE
	subject := "Reset Password from LocalEats"

	to := []string{
		email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + code + "\r\n")

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	err = rdb.Set(context.Background(), "bahriddinmamarajabov94@gmail.com", code, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to save verification code to Redis: %v", err)
	}

	return nil
}

// bahriddinmamarajabov94@gmail.com
