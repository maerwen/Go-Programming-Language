package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 {
	if username == "joe@example.org" {
		return 980000000
	}
	return 0
	/*  */
}

const sender = "notifications@example.com"
const password = "correntPassword"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage,%d%% of your quota.`

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 //1GB
	percent := 100 * used / quota
	if percent < 90 {
		return //ok
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", username, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}
