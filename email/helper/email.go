package helper

import (
	"bytes"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
)

type SendEmailParams struct {
	To      string
	Subject string
	Code    string
	Title   string
	Time    int32
}

func SendEmail(params SendEmailParams) error {
	tmpl, err := template.ParseFiles("./template/email-code.html")
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)

	err = tmpl.Execute(body, map[string]interface{}{
		"Code":  params.Code,
		"Title": params.Title,
		"Time":  params.Time,
	})

	if err != nil {
		return err
	}

	e := &email.Email{
		From:    "1761617270@qq.com",
		To:      []string{params.To},
		Subject: params.Subject,
		HTML:    body.Bytes(),
	}

	_ = e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1761617270@qq.com", "giqzyspehhdddaig", "smtp.qq.com"))
	
	return nil
}
