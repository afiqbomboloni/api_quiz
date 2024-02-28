package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type Config struct {
	DBDriver string
	DBHost   string
	DBPass   string
	DBName   string
	DBUser   string
	RedisAddr	string
	RedisPass	string
	DBPort   int
}

func LoadConfig() (Config) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file failed")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("DB_PORT failed")
	
	}

	return Config{
		DBDriver: os.Getenv("DB_DRIVER"),
		DBHost: os.Getenv("DB_HOST"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		RedisAddr: os.Getenv("REDIS_ADDRESS"),
		RedisPass: os.Getenv("REDIS_PASSWORD"),
		DBPort: dbPort,
	}
}


func ConnectDb() *gorm.DB{

	cfg := LoadConfig()
	
	// "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db

}