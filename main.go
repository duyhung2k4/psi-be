package main

import (
	"log"
	"net/http"
	"psi/config"
	_ "psi/docs"
	"psi/router"
	"psi/utils"
	"time"
)

// @title psi
// @version 1.0
// @description psi

// @host localhost:18888
// @BasePath /api/v1
func main() {
	server := &http.Server{
		Addr:         "localhost:" + config.GetAppPort(),
		Handler:      router.Router(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	utils.Cronjob()

	if err := server.ListenAndServe(); err != nil {
		log.Println("Error run server: ", err)
	}

}

// package main

// import (
// 	"fmt"
// 	"net/smtp"
// )

// func main() {
// 	// Sender data.
// 	from := "nguyenduyhung04092004@gmail.com"
// 	password := "sogxzutksoipfhnj"

// 	// Receiver email address.
// 	to := []string{
// 		"baba0398262124@gmail.com",
// 	}

// 	// smtp server configuration.
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	// Message.
// 	message := []byte("This is a test email message.")

// 	// Authentication.
// 	auth := smtp.PlainAuth("", from, password, smtpHost)

// 	// Sending email.
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Email Sent Successfully!")
// }
