package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct{
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb(){
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil{
		log.Fatal("Error al conectar con la base de datos! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Conectado con exito")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Arrancar Migrations")





}