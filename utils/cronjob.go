package utils

import (
	"log"
	"psi/cronjob"

	"github.com/robfig/cron/v3"
)

func Cronjob() {
	s := cron.New()

	s.AddFunc("@every 600s", cronjob.DeleteExpiredTemporaryEmail)

	log.Println("Start cron job")
	s.Start()
}
