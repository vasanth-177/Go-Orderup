package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func sendMail() {

	m := gomail.NewMessage()
	m.SetHeader("From", "vasanthkai17@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Order Up!!!")
	m.SetBody("text/html", "Order details attached....\nDelivered within 30mins.....")
	m.Attach("pdfs/" + uname + ".pdf")

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "vasanthkai17@gmail.com", "qvremjzkjqztgbii")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("Mail Sent.....")
}
