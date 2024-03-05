package main

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/afiqbomboloni/api_quiz/config"
	"github.com/afiqbomboloni/api_quiz/entity"

)

func main() {
	config.LoadConfig()


	db := config.ConnectDb()

	data := createAdminUser(db)
	if data != "" {
		fmt.Println(data)
	}

}


func createAdminUser(db *gorm.DB) string {
	

	
	if err := db.WithContext(context.Background()).
		Model(&entity.User{}).
		Create(entity.NewUser(
			"admin yes4",
			"admin4@gmail.com",
			"testing1234",
			"admin",
		)).
		Error; err != nil {
		panic(err)
	} 

	return "admin user created"
}

