package config

import (
	"fmt"
	"log"
	"psi/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect(migrate bool) (*gorm.DB, error) {
	dns := fmt.Sprintf(
		`	host=%s
			user=%s
			password=%s
			dbname=%s
			port=%s
			sslmode=disable`,
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	var err error

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if migrate {
		log.Println("Start migrate")
		err := db.AutoMigrate(
			&model.Profile{},
			&model.Credential{},
			&model.Role{},
			&model.TemporaryCredential{},

			&model.Course{},
			&model.Subject12{},
			&model.Subject12Course{},
			&model.Major{},
			&model.MajorCourse{},
			&model.LanguageCertificate{},
			&model.LanguageCertificateCourse{},
			&model.Skill{},
			&model.SkillCourse{},
			&model.SaveCourse{},
			&model.RegisterCourse{},
		)

		if err != nil {
			return nil, err
		}

		log.Println("Migrate done!")
	}

	return db, nil
}
