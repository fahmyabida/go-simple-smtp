package main

import (
	"fmt"
	"log"
	"net/smtp"
)

// sendEmail is ...
func sendEmail() {
	server := "smtp-mail.outlook.com"
	port := 587
	user := "rolandbrilianto@outlook.com"
	from := user
	pass := "Kipasangin_1"
	
	// user := "your outlook email"
	// from := user
	// pass := "your password"
	dest := "fahmyabida@gmail.com"

	auth := LoginAuth(user, pass)

	to := []string{dest}

	msg := []byte("From: " + from + "\n" +
		"To: " + dest + "\n" +
		"Subject: Test outlook\n\n" +
		"OK this is message")

	endpoint := fmt.Sprintf("%v:%v", server, port)
	err := smtp.SendMail(endpoint, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
