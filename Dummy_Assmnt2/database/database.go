package database

import (
	//"fmt"
	"fmt"
	"log"
	"os"

	"example/Dummy_Assmnt2/models"
	//"gorm.io/driver/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	///
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {

	// dsn := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=mydatabase sslmode=disable"
	// a := os.Getenv("DB_PASSWORD")
	//fmt.Println("-------------------")
	// fmt.Println(a)
	//str, _ := os.LookupEnv("DB_HOST")
	//fmt.Printf("str: %v\n", str)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		// os.LookupEnv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// fmt.Println(dsn)
	fmt.Printf("dsn: %v\n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{
		Db: db,
	}
}
