package app

import (
	// "github.com/afiqbomboloni/api_quiz/middleware"
	"github.com/afiqbomboloni/api_quiz/middleware"
	authHandlerv1 "github.com/afiqbomboloni/api_quiz/modules/auth/v1/handler"
	authServicev1 "github.com/afiqbomboloni/api_quiz/modules/auth/v1/service"
	jawabanPesertaHandlerv1 "github.com/afiqbomboloni/api_quiz/modules/jawabanPeserta/v1/handler"
	jawabanPesertaServicev1 "github.com/afiqbomboloni/api_quiz/modules/jawabanPeserta/v1/service"
	pertanyaanHandlerv1 "github.com/afiqbomboloni/api_quiz/modules/pertanyaan/v1/handler"
	pertanyaanServicev1 "github.com/afiqbomboloni/api_quiz/modules/pertanyaan/v1/service"
	quizHandlerv1 "github.com/afiqbomboloni/api_quiz/modules/quiz/v1/handler"
	quizServicev1 "github.com/afiqbomboloni/api_quiz/modules/quiz/v1/service"

	"github.com/gin-gonic/gin"
)

func QuizHTTPHandler(router *gin.Engine, as quizServicev1.QuizService) {
	h := quizHandlerv1.NewQuizHandler(as)
	v1 := router.Group("v1")

	v1.GET("/quizzes", h.GetQuizzes)
	v1.GET("/quizzes/:id", h.GetQuiz)
	v1.Use(middleware.AuthMiddleware())
	{
		
		v1.POST("/quizzes", h.CreateQuiz)
		v1.PUT("/quizzes/:id", h.UpdateQuiz)
		v1.DELETE("/quizzes/:id", h.DeleteQuiz)
	}
	
}


func PertanyaanHTTPHandler(router *gin.Engine, ps pertanyaanServicev1.PertanyaanService) {
	h := pertanyaanHandlerv1.NewPertanyaanHandler(ps)
	v1 := router.Group("v1")

	v1.GET("/questions", h.GetPertanyaans)
	v1.GET("/questions/:id", h.GetPertanyaan)
	v1.POST("/questions", h.CreatePertanyaan)
	v1.PUT("/questions/:id", h.UpdatePertanyaan)
	v1.DELETE("/questions/:id", h.DeletePertanyaan)
	v1.GET("/questions/quiz/:id", h.GetPertanyaanByQuiz)
	// v1.Use(middleware.AuthMiddleware())
	// {
		
	// }
	
}

func JawabanPesertaHTTPHandler(router *gin.Engine, jps jawabanPesertaServicev1.JawabanPesertaService) {
	h := jawabanPesertaHandlerv1.NewJawabanPesertaHandler(jps)
	v1 := router.Group("v1")

	v1.GET("/answers", h.GetJawabanPesertas)
	v1.GET("/answers-quiz", h.GetJawabanByUserAndQuiz)
	v1.GET("/answer-users", h.ListResultQuiz)
	v1.PUT("/answers/skor", h.UpdateSkor)
	v1.GET("/answers/:id", h.GetJawabanPeserta)
	v1.POST("/answers", h.CreateJawabanPeserta)
	v1.PUT("/answers/:id", h.UpdateJawabanPeserta)
	v1.DELETE("/answers/:id", h.DeleteJawabanPeserta)
	// v1.Use(middleware.AuthMiddleware())
	// {
		
	// }
	
}



func AuthHTTPHandler(router *gin.Engine, as authServicev1.AuthService) {
	h :=authHandlerv1.NewAuthHandler(as)
	v1 := router.Group("v1")

	
	v1.POST("/register", h.Register)
	v1.POST("/login", h.Login)
	v1.POST("/logout", h.Logout)
}