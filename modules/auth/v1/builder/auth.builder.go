package builder

import (
	"github.com/afiqbomboloni/api_quiz/app"
	"github.com/afiqbomboloni/api_quiz/config"

	"github.com/afiqbomboloni/api_quiz/modules/auth/v1/repository"
	"github.com/afiqbomboloni/api_quiz/modules/auth/v1/service"

	"github.com/gin-gonic/gin"
)
func AuthorBuilder(router *gin.Engine) {
	repo := repository.NewAuthRepository(config.ConnectDb())
    svc := service.NewAuthRepository(repo)
	
	app.AuthHTTPHandler(router, svc)
}