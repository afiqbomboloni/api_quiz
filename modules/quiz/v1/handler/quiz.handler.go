package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/afiqbomboloni/api_quiz/modules/quiz/v1/service"
	"github.com/afiqbomboloni/api_quiz/request"
	"github.com/afiqbomboloni/api_quiz/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type quizHandler struct {
	quizService service.QuizService
}

func NewQuizHandler(quizService service.QuizService) *quizHandler {
	return &quizHandler{quizService}
}

// pagination for quiz with default limit 10 and page 1
func (h *quizHandler) GetQuizzes(ctx *gin.Context) {
	limitString := ctx.DefaultQuery("limit", "5")
	pageString := ctx.DefaultQuery("page", "1")
	isNotExpired := ctx.DefaultQuery("isNotExpired", "true")

	limit, err := strconv.ParseUint(limitString, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
        return
    }

    page, err := strconv.ParseUint(pageString, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
        return
    }
	
	quizzes, err := h.quizService.FindAll(limit,page, isNotExpired)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	

	var quizzesResponse []response.QuizResponse
	for _, quiz := range quizzes {
		total_soal,_ := h.quizService.CountQuestion(quiz.ID)
		quizzesResponse = append(quizzesResponse, response.NewQuizResponse(quiz, total_soal))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    quizzesResponse,
	})
}


func (h *quizHandler) GetQuiz(ctx *gin.Context) {
	idString := ctx.Param("id")

	idUint, _ := strconv.ParseUint(idString, 10, 64)
	id := uint(idUint)

	quiz, err := h.quizService.FindByID(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "not found",
			"error":   err,
		})
		return
	}

	quizResponse := response.QuizResponse{
		ID:           quiz.ID,
		Judul:        quiz.Judul,
		Deskripsi:    quiz.Deskripsi,
		WaktuMulai:   quiz.WaktuMulai.Format("2006-01-02T15:04:05Z"),
		WaktuSelesai: quiz.WaktuSelesai.Format("2006-01-02T15:04:05Z"),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    quizResponse,
	})
}

func (h *quizHandler) CreateQuiz(ctx *gin.Context) {
    var quizRequest request.QuizRequest
    err := ctx.ShouldBindJSON(&quizRequest)
    if err != nil {
        errorMessages := []string{}
        for _, e := range err.(validator.ValidationErrors) {
            errorMessage := fmt.Sprintf("Error on field %s, is %s", e.Field(), e.ActualTag())
            errorMessages = append(errorMessages, errorMessage)
        }
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "Bad Request",
            "error":   errorMessages,
        })
        return
    }

    h.quizService.Create(quizRequest)
    ctx.JSON(http.StatusCreated, gin.H{
        "message": "new quiz created",
    })
}

func (h *quizHandler) UpdateQuiz(ctx *gin.Context) {
	var quizRequest request.QuizRequest
	err := ctx.ShouldBindJSON(&quizRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, is %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   errorMessages,
		})
		return
	}

	idString := ctx.Param("id")
	idUint, _ := strconv.ParseUint(idString, 10, 64)
	id := uint(idUint)

	h.quizService.Update(id, quizRequest)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "quiz updated",
	},
)

}

func (h *quizHandler) DeleteQuiz(ctx *gin.Context) {
	idString := ctx.Param("id")
	idUint, _ := strconv.ParseUint(idString, 10, 64)
	id := uint(idUint)

	h.quizService.Delete(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "quiz deleted",
	},
	)
}
