package main

import (
	"echo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=alghiffaryfa19 password=alghiffaryfa19 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(
		&models.User{},
		&models.Mentor{},
		&models.Course{},
		&models.Video{},
		&models.Attachment{},
	)

}