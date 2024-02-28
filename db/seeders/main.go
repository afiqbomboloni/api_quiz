package main

import (
	"context"
	"fmt"

	// "github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/afiqbomboloni/api_quiz/config"
	"github.com/afiqbomboloni/api_quiz/entity"
	// "github.com/afiqbomboloni/api_quiz/utils"
)

func main() {
	config.LoadConfig()
	// checkError(err)

	db := config.ConnectDb()
	// checkError(err)

	// if cfg.Env == "production" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	// createSampleUser(db)
	data := createAdminUser(db)
	if data != "" {
		fmt.Println(data)
	}

}

// func createSampleUser(db *gorm.DB) {
// 	userID := uuid.New()

// 	dob, _ := utils.DateStringToTime("1996-11-04")

// 	if err := db.WithContext(context.Background()).
// 		Model(&entity.User{}).
// 		Create(entity.NewUser(
// 			userID,
// 			"Bayu Novianto",
// 			"bayunoviantoo9@gmail.com",
// 			"test123",
// 			utils.TimeToNullTime(dob),
// 			"",
// 			"0895346419497",
// 			"system",
// 		)).
// 		Error; err != nil {
// 		panic(err)
// 	}

// 	userID2 := uuid.New()

// 	dob2, _ := utils.DateStringToTime("1996-11-04")

// 	if err := db.WithContext(context.Background()).
// 		Model(&entity.User{}).
// 		Create(entity.NewUser(
// 			userID2,
// 			"User Test",
// 			"user-test@gmail.com",
// 			"testingApp23!",
// 			utils.TimeToNullTime(dob2),
// 			"",
// 			"",
// 			"system",
// 		)).
// 		Error; err != nil {
// 		panic(err)
// 	}
// }

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

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }