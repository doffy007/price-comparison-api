package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Setup database connection
func SetUpDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("cannot load on database, try again or check your programs")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cannot load on database, try again or check your programs")
	}
	// model DB
	// db.AutoMigrate()
	return db
}

//Close connection database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("cannot load on database, try again or check your programs")
	}
	dbSQL.Close()
}