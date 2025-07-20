package services

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "log"
    "net/smtp"
    "os"
)

type EmailService struct {
    SMTPHost     string
    SMTPPort     string
    SMTPUsername string
    SMTPPassword string
    FromEmail    string
}

func NewEmailService() *EmailService {
    return &EmailService{
        SMTPHost:     os.Getenv("SMTP_HOST"),
        SMTPPort:     os.Getenv("SMTP_PORT"),
        SMTPUsername: os.Getenv("SMTP_USERNAME"),
        SMTPPassword: os.Getenv("SMTP_PASSWORD"),
        FromEmail:    os.Getenv("FROM_EMAIL"),
    }
}

func (e *EmailService) GenerateVerificationToken() (string, error) {
    bytes := make([]byte, 32)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

func (e *EmailService) SendVerificationEmail(email, token, userName string) error {
	log.Printf("[DEBUG] Iniciando envío de email a: %s", email)
	log.Printf("[DEBUG] SMTP Config - Host: %s, Port: %s, Username: %s", 
		e.SMTPHost, e.SMTPPort, e.SMTPUsername)
	
	if e.SMTPHost == "" || e.SMTPPort == "" || e.SMTPUsername == "" || e.SMTPPassword == "" {
		log.Printf("[ERROR] Configuración SMTP incompleta")
		log.Printf("[ERROR] Host: '%s', Port: '%s', Username: '%s', Password: '%s'", 
			e.SMTPHost, e.SMTPPort, e.SMTPUsername, "[OCULTO]")
		return fmt.Errorf("configuración SMTP incompleta")
	}
	
	verificationURL := fmt.Sprintf("http://localhost:8080/api/auth/verify-email?token=%s", token)
	log.Printf("[DEBUG] URL de verificación: %s", verificationURL)
	
	subject := "Confirma tu cuenta - Menchaca Health"
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>¡Hola %s!</h2>
			<p>Gracias por registrarte en Menchaca Health.</p>
			<p>Para activar tu cuenta, haz clic en el siguiente enlace:</p>
			<a href="%s" style="background-color: #4CAF50; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px;">Verificar Email</a>
			<p>Este enlace expira en 24 horas.</p>
			<p>Si no te registraste en nuestra plataforma, ignora este email.</p>
		</body>
		</html>
	`, userName, verificationURL)

	// Construir el mensaje completo
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s", email, subject, body)

	log.Println("[DEBUG] Configurando autenticación SMTP")
	// Configurar autenticación
	auth := smtp.PlainAuth("", e.SMTPUsername, e.SMTPPassword, e.SMTPHost)
	addr := fmt.Sprintf("%s:%s", e.SMTPHost, e.SMTPPort)
	
	log.Printf("[DEBUG] Intentando conectar a: %s", addr)
	
	// Enviar el email
	err := smtp.SendMail(addr, auth, e.FromEmail, []string{email}, []byte(message))
	if err != nil {
		log.Printf("[ERROR] Error enviando email: %v", err)
		return fmt.Errorf("error enviando email: %v", err)
	}
	
	log.Printf("[DEBUG] Email enviado exitosamente a: %s", email)
	return nil
}