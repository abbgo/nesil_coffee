package controllers

import (
	"bytes"
	"context"
	"html/template"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

var authh smtp.Auth

func SendMail(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var mail models.ForMail
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
		Name         string
		Mail         string
		Message      string
		ProductName  string
		ProductImage string
	}{
		Name:         mail.FullName,
		Mail:         mail.Email,
		Message:      mail.Letter,
		ProductName:  "",
		ProductImage: "",
	}

	parsedTemplate := "templates/template.html"

	if mail.ProductID.String != "" {
		var productName, productImage string
		if err := db.QueryRow(context.Background(), `SELECT name_tm FROM products WHERE id=$1`, mail.ProductID).
			Scan(&productName); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if err := db.QueryRow(context.Background(), `SELECT image FROM product_images WHERE product_id=$1 LIMIT 1`, mail.ProductID).
			Scan(&productImage); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		templateData = struct {
			Name         string
			Mail         string
			Message      string
			ProductName  string
			ProductImage string
		}{
			Name:         mail.FullName,
			Mail:         mail.Email,
			Message:      mail.Letter,
			ProductName:  productName,
			ProductImage: productImage,
		}

		parsedTemplate = "templates/template_product.html"
	}

	r := NewRequest(to, "Salam. Men "+templateData.Name, "Salam Nesil Coffee !")
	if err := r.ParseTemplate(serverPath+parsedTemplate, templateData); err == nil {
		_, err := r.SendEmail()
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	} else {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var productID interface{}
	if mail.ProductID.String == "" {
		productID = nil
	} else {
		productID = mail.ProductID
	}

	_, err = db.Exec(context.Background(), `INSERT INTO mails (full_name,email,letter,product_id) VALUES ($1,$2,$3,$4)`,
		mail.FullName, mail.Email, mail.Letter, productID)
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
