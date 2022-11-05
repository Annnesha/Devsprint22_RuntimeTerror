package database

import (
  "os"

  "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"  
)

var DB *gorm.DB

func Connect(){

  mySecretDialect := os.Getenv("DIALECT")
  mySecretURI := os.Getenv("DB_URI")

  connect, gorm_err := gorm.Open(mySecretDialect, mySecretURI)

  if gorm_err != nil {
    panic("Unable to connect")

  }

  DB = connect

  connect.AutoMigrate(&models.Student{}, &models.Teacher{}, &models.Events{})
}