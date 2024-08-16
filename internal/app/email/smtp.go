package email

import (
	"crypto/tls"
	"fmt"

	"github.com/aparnasukesh/notification-svc/config"
	gomail "gopkg.in/mail.v2"
)

type smtpEmail struct {
	email    string
	password string
}

type SmtpEmail interface {
	SendEmail(otp, emails string) error
}

func NewSMTPEmail(cfg config.Config) SmtpEmail {
	return &smtpEmail{
		email:    cfg.EMAIL,
		password: cfg.PASSWORD,
	}
}

func (s smtpEmail) SendEmail(otp, emails string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", s.email)

	m.SetHeader("To", emails)

	m.SetHeader("Subject", "OTP to verify your Gmail")

	m.SetBody("text/plain", otp+" is your OTP to register to ShoeZone. Thank you registering to our site. Dont't give this code to anyone")

	d := gomail.NewDialer("smtp.gmail.com", 587, s.email, s.password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	fmt.Println("OTP has been sent successfully")
	return nil
}
