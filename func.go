package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/mail"
	"net/smtp"

	fdk "github.com/fnproject/fdk-go"
)

type payload struct {
	To      string
	Subject string
	Body    string
}

func main() {
	fdk.Handle(fdk.HandlerFunc(func(ctx context.Context, in io.Reader, out io.Writer) {
		var p payload
		err := json.NewDecoder(in).Decode(&p)
		if err != nil {
			log.Println("failed to parse email", err.Error())
			out.Write([]byte("failed to parse email " + err.Error()))
			return
		}

		to, err := mail.ParseAddress(p.To)
		if err != nil {
			log.Println("failed to parse address", err.Error())
			out.Write([]byte("failed to parse address " + err.Error()))
			return
		}

		fnCtx := fdk.GetContext(ctx)
		cfg := fnCtx.Config()
		username := cfg["OCI_EMAIL_DELIVERY_USER_OCID"]
		password := cfg["OCI_EMAIL_DELIVERY_USER_PASSWORD"]
		server := cfg["OCI_EMAIL_DELIVERY_SMTP_SERVER"]
		sender := cfg["OCI_EMAIL_DELIVERY_APPROVED_SENDER"]

		log.Println("OCI_EMAIL_DELIVERY_USER_OCID", username)
		log.Println("OCI_EMAIL_DELIVERY_USER_PASSWORD", password)
		log.Println("OCI_EMAIL_DELIVERY_SMTP_SERVER", server)
		log.Println("OCI_EMAIL_DELIVERY_APPROVED_SENDER", sender)

		from, err := mail.ParseAddress(sender)
		if err != nil {
			log.Println("failed to parse address", err.Error())
			out.Write([]byte("failed to parse address " + err.Error()))
			return
		}

		message := []byte("" +
			"From: " + from.String() + "\r\n" +
			"To: " + to.String() + "\r\n" +
			"Subject: " + p.Subject + "\r\n" +
			"\r\n" +
			p.Body + "\r\n")

		err = smtp.SendMail(server+":25", smtp.PlainAuth("", username, password, server), from.Address, []string{to.Address}, message)
		if err != nil {
			log.Println("failed to send email", err.Error())
			out.Write([]byte("failed to send email " + err.Error()))
			return
		}

		log.Println("sent email successfully!")
		out.Write([]byte("sent email successfully!"))
	}))
}
