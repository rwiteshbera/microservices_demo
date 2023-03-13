package handlers

import (
	"bytes"
	"html/template"

	"github.com/vanng822/go-premailer/premailer"
)

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type Message struct {
	FromMail    string
	FromName    string
	ToMail      string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
}

func (m *Mail) SendSMTPMessage(message Message) error {
	if message.FromMail == "" {
		message.FromMail = message.FromMail
	}
	if message.FromName == "" {
		message.FromName = message.FromName
	}

	data := map[string]any{
		"message": message.Data,
	}

	message.DataMap = data

	formattedMessageStructure, err := m.buildHTMLMessage(message)
	if err != nil {
		return err
	}

	plainMessage, err := m.buildPlainTextMessage(message)
	if err != nil {
		return err
	}

}

func (m *Mail) buildHTMLMessage(messsage Message) (string, error) {
	templateToRender := "../templates/mail.html.gohtml"

	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tmplt bytes.Buffer
	if err = t.ExecuteTemplate(&tmplt, "body", messsage.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tmplt.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}
	return formattedMessage, nil
}

func (m *Mail) buildPlainTextMessage(messsage Message) (string, error) {
	templateToRender := "../templates/mail.plain.gohtml"

	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tmplt bytes.Buffer
	if err = t.ExecuteTemplate(&tmplt, "body", messsage.DataMap); err != nil {
		return "", err
	}

	plainMessage := tmplt.String()

	return plainMessage, nil
}

func (m *Mail) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}
	return html, nil
}
