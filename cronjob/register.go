package cronjob

import (
	"psi/config"
	"psi/model"
	"time"
)

func DeleteExpiredTemporaryEmail() {
	db := config.GetDB()

	db.Where("time_end < ?", time.Now()).Unscoped().Delete(&model.TemporaryCredential{})
}
