package service

import (
	"fmt"
	"net/smtp"

	"windsurf-project/internal/config"
)

type EmailService struct {
	cfg *config.Config
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{cfg: cfg}
}

func (s *EmailService) SendPasswordResetEmail(email, token string) error {
	if s.cfg.SMTPUser == "" || s.cfg.SMTPPassword == "" {
		// In development, just log the token
		fmt.Printf("\n=== PASSWORD RESET TOKEN ===\n")
		fmt.Printf("Email: %s\n", email)
		fmt.Printf("Token: %s\n", token)
		fmt.Printf("Reset URL: %s/reset-password?token=%s\n", s.cfg.FrontendURL, token)
		fmt.Printf("===========================\n\n")
		return nil
	}

	resetURL := fmt.Sprintf("%s/reset-password?token=%s", s.cfg.FrontendURL, token)
	
	subject := "Password Reset Request"
	body := fmt.Sprintf(`
Hello,

You have requested to reset your password. Please click the link below to reset your password:

%s

This link will expire in 1 hour.

If you did not request this password reset, please ignore this email.

Best regards,
Social App Team
`, resetURL)

	message := fmt.Sprintf("From: %s\r\n", s.cfg.SMTPUser)
	message += fmt.Sprintf("To: %s\r\n", email)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += "\r\n" + body

	auth := smtp.PlainAuth("", s.cfg.SMTPUser, s.cfg.SMTPPassword, s.cfg.SMTPHost)
	addr := fmt.Sprintf("%s:%s", s.cfg.SMTPHost, s.cfg.SMTPPort)

	err := smtp.SendMail(addr, auth, s.cfg.SMTPUser, []string{email}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (s *EmailService) SendWelcomeEmail(email, username string) error {
	if s.cfg.SMTPUser == "" || s.cfg.SMTPPassword == "" {
		// In development, just log
		fmt.Printf("\n=== WELCOME EMAIL ===\n")
		fmt.Printf("Email: %s\n", email)
		fmt.Printf("Username: %s\n", username)
		fmt.Printf("====================\n\n")
		return nil
	}

	subject := "Welcome to Social App!"
	body := fmt.Sprintf(`
Hello %s,

Welcome to Social App! We're excited to have you join our community of foreigners connecting through shared interests.

Explore our interest groups:
- Coworking: Connect with professionals and digital nomads
- Photography: Share and discuss your photography
- Food: Discover local cuisine and restaurants
- Languages: Practice and learn new languages

Get started by completing your profile and joining your first interest group!

Best regards,
Social App Team
`, username)

	message := fmt.Sprintf("From: %s\r\n", s.cfg.SMTPUser)
	message += fmt.Sprintf("To: %s\r\n", email)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += "\r\n" + body

	auth := smtp.PlainAuth("", s.cfg.SMTPUser, s.cfg.SMTPPassword, s.cfg.SMTPHost)
	addr := fmt.Sprintf("%s:%s", s.cfg.SMTPHost, s.cfg.SMTPPort)

	err := smtp.SendMail(addr, auth, s.cfg.SMTPUser, []string{email}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
