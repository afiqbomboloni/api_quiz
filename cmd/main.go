package main

import (
	"github.com/afiqbomboloni/api_quiz/config"
	"github.com/afiqbomboloni/api_quiz/entity"
	"github.com/afiqbomboloni/api_quiz/middleware"
	authBuilder "github.com/afiqbomboloni/api_quiz/modules/auth/v1/builder"
	jawabanPesertaBuilder "github.com/afiqbomboloni/api_quiz/modules/jawabanPeserta/v1/builder"
	pertanyaanBuilder "github.com/afiqbomboloni/api_quiz/modules/pertanyaan/v1/builder"
	quizBuilder "github.com/afiqbomboloni/api_quiz/modules/quiz/v1/builder"


	"fmt"
	

	"github.com/gin-gonic/gin"
)

func main() {


	config.LoadConfig()

	
	config.ConnectDb().Debug().AutoMigrate(&entity.Quiz{}, &entity.Pertanyaan{}, &entity.JawabanPeserta{}, &entity.User{})
	fmt.Println("Success")

	r := gin.New()
	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.LoggerMiddleware())
	

	BuildHandler(r)


	r.Run("localhost:8080")
}

func BuildHandler(router *gin.Engine) {
	quizBuilder.QuizBuilder(router)
	pertanyaanBuilder.PertanyaanBuilder(router)
	jawabanPesertaBuilder.JawabanPesertaBuilder(router)
	authBuilder.AuthorBuilder(router)



}

