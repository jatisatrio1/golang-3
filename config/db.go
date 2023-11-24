package config

import (
	"fmt"
	"weather/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db-go-sql"
)

var (
	db  *gorm.DB
	err error
)

func connect() {
	psqlInfo := fmt.Sprintf(`
	host=%s 
	port=%d 
	user=%s `+`
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, password, dbname)

	//ini koneksi menggunakan database/sql
	// db, err = sql.Open("postgres", psqlInfo)

	//ini koneksi menggunakan gorm
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	db.Debug().AutoMigrate(&model.Weather{})
}

func GetDB() *gorm.DB {
	return db
}
