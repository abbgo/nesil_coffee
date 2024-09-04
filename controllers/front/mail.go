package controllers

import (
	"bytes"
	"context"
	"html/template"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

var authh smtp.Auth

type ForMail struct {
	FullName string `json:"full_name" binding:"required,min=3"`
	Email    string `json:"email" binding:"email"`
	Letter   string `json:"letter" binding:"required,min=3"`
}

func SendMail(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var mail ForMail
	if err := c.BindJSON(&mail); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var from = os.Getenv("MAIL_FROM")
	var password = os.Getenv("MAIL_PASSWORD")
	var toEmailAddress = os.Getenv("MAIL_TO")
	var to = []string{toEmailAddress}

	var host = os.Getenv("MAIL_HOST")
	var serverPath = os.Getenv("SERVER_PATH")

	authh = smtp.PlainAuth("", from, password, host)
	templateData := struct {
		Name    string
		Mail    string
		Message string
	}{
		Name:    mail.FullName,
		Mail:    mail.Email,
		Message: mail.Letter,
	}
	r := NewRequest(to, "Salam. Men "+templateData.Name, "Salam Nesil Coffee !")
	if err := r.ParseTemplate(serverPath+"templates/template.html", templateData); err == nil {
		_, err := r.SendEmail()
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	} else {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), `INSERT INTO mails (full_name,mail,letter) VALUES ($1,$2,$3)`,
		mail.FullName, mail.Email, mail.Letter)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "mail successfully send",
	})

}

// Request struct
type Request struct {
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	var port = os.Getenv("MAIL_PORT")
	var from = os.Getenv("MAIL_FROM")
	var host = os.Getenv("MAIL_HOST")

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := host + ":" + port

	if err := smtp.SendMail(addr, authh, from, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
