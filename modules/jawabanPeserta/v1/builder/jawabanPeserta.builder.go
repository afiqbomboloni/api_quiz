package builder

import (
	"github.com/afiqbomboloni/api_quiz/app"
	"github.com/afiqbomboloni/api_quiz/config"
	"github.com/gin-gonic/gin"
	"github.com/afiqbomboloni/api_quiz/modules/jawabanPeserta/v1/repository"
	"github.com/afiqbomboloni/api_quiz/modules/jawabanPeserta/v1/service"
)

func JawabanPesertaBuilder(router *gin.Engine) {

	repo := repository.NewJawabanPesertaRepository(config.ConnectDb())
    svc := service.NewJawabanPesertaService(repo)
	
	app.JawabanPesertaHTTPHandler(router, svc)
}