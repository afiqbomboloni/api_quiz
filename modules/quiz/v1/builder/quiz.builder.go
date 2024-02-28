package builder

import (
	"github.com/afiqbomboloni/api_quiz/app"
	"github.com/afiqbomboloni/api_quiz/config"
	// "github.com/afiqbomboloni/api_quiz/utils"
	"github.com/afiqbomboloni/api_quiz/modules/quiz/v1/repository"
	"github.com/afiqbomboloni/api_quiz/modules/quiz/v1/service"
	"github.com/gin-gonic/gin"

)
func QuizBuilder(router *gin.Engine) {

	repo := repository.NewQuizRepository(config.ConnectDb())
    svc := service.NewQuizService(repo)
	
	app.QuizHTTPHandler(router, svc)
}