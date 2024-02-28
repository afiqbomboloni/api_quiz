package builder

import (
	"github.com/afiqbomboloni/api_quiz/app"
	"github.com/afiqbomboloni/api_quiz/config"
	"github.com/afiqbomboloni/api_quiz/modules/pertanyaan/v1/repository"
	"github.com/afiqbomboloni/api_quiz/modules/pertanyaan/v1/service"
	"github.com/gin-gonic/gin"

)
func PertanyaanBuilder(router *gin.Engine) {

	repo := repository.NewPertanyaanRepository(config.ConnectDb())
    svc := service.NewPertanyaanService(repo)
	
	app.PertanyaanHTTPHandler(router, svc)
}