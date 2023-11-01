package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

func main() {
	username := "ahrorzoda111@gmail.com"
	password := "fhea cagf mwxv xunb"
	m := gomail.NewMessage()
	m.SetHeader("From", "Shohrukh Ahrorzoda <"+username+">")
	m.SetHeader("To", "firdavs080902@gmail.com")
	m.SetHeader("Subject", "бобой канд")
	m.SetBody("text/plain", "18+")
	d := gomail.NewDialer("smtp.gmail.com", 587, username, password)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}
