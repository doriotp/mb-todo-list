package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	gomail "gopkg.in/mail.v2"
)

var secretKey = []byte("secretpassword")

// GenerateToken generates a JWT token with the user ID as part of the claims
func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token valid for 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for errors
	if err != nil {
		return nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func SendEmail(token, recipientEmail string)error{
	var (
		baseURL = "https://example.com/reset-password"
	resetLink = fmt.Sprintf("%s?token=%s", baseURL, token)
	from = "amitkhan1911@gmail.com"
	)

	// Create a new message
	message := gomail.NewMessage()

	// Set email headers
	message.SetHeader("From", from)
	message.SetHeader("To", recipientEmail)
	message.SetHeader("Subject", "Password Reset Link")

	// Set email body with the reset password link
	body := fmt.Sprintf("Hello,\n\nYou requested a password reset. Please click the link below to reset your password:\n\n%s\n\nIf you did not request this, please ignore this email.\n\nThanks,\nYour Team", resetLink)
	message.SetBody("text/plain", body)

	// Set up the SMTP dialer
	dialer := gomail.NewDialer("smtp.gmail.com", 587, from, "fjev ijgi isnw oplb")

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}


