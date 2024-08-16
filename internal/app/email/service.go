package email

type service struct {
	smtp SmtpEmail
}

type Service interface {
	SendEmail(otp, emails string) error
}

func NewService(smtp SmtpEmail) Service {
	return &service{
		smtp: smtp,
	}
}

func (s service) SendEmail(otp, email string) error {
	return s.smtp.SendEmail(otp, email)
}
