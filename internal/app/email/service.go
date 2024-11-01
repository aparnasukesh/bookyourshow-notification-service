package email

type service struct {
	smtp SmtpEmail
}

type Service interface {
	SendEmail(otp, emails string) error
	SendResetPassWordEmail(otp, email string) error
}

func NewService(smtp SmtpEmail) Service {
	return &service{
		smtp: smtp,
	}
}

func (s service) SendEmail(otp, email string) error {
	return s.smtp.SendEmail(otp, email)
}

func (s *service) SendResetPassWordEmail(otp, email string) error {
	return s.smtp.SendResetPassWordEmail(otp, email)
}
