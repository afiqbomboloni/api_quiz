package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/afiqbomboloni/api_quiz/modules/jawabanPeserta/v1/service"
	"github.com/afiqbomboloni/api_quiz/request"
	"github.com/afiqbomboloni/api_quiz/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jwabanPesertaHandler struct {
	jawabanPesertaService service.JawabanPesertaService
}

func NewJawabanPesertaHandler(jawabanPesertaService service.JawabanPesertaService) *jwabanPesertaHandler {
	return &jwabanPesertaHandler{jawabanPesertaService}
}

func(h *jwabanPesertaHandler) ListResultQuiz(ctx *gin.Context) {
	// ini adalah list dengan array dan response seperti jawabanpesertaadminresponse akan tetapi id_usernya unik
	jawaban_peserta, err := h.jawabanPesertaService.FindAllWIthUniqueUser()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	var jawaban_pesertaResponse []response.JawabanPesertaAdminResponse
	for _, p := range jawaban_peserta {
	
		total_soal,_ := h.jawabanPesertaService.CountQuestion(p.IdQuiz)
		jumlah_benar := h.jawabanPesertaService.CountRightAnswer(p.IdQuiz, p.IdUser, p.IdPertanyaan, p.JawabanPeserta)
		jawaban_pesertaResponse = append(jawaban_pesertaResponse, response.NewJawabanPesertaAdminResponse(p, total_soal, jumlah_benar))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    jawaban_pesertaResponse,
	})
}

// updateSkor dengan request skor, id_user dan param id jawaban peserta. akan mengupdate skor jawaban peserta
// misal input adalah 90 dan akan mencari semua id user yang memiliki jawaban benar dan mengupdate skornya menjadi 90/total jawaban benar
func(h *jwabanPesertaHandler) UpdateSkor(ctx *gin.Context) {
	var jawaban_pesertaRequest request.UpdateSkorRequest
	fmt.Println(jawaban_pesertaRequest)
	err := ctx.ShouldBindJSON(&jawaban_pesertaRequest)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
		
			errorMessages := []string{}
			for _, e := range validationErrors {
				errorMessage := fmt.Sprintf("Error on field %s, is %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
				"error":   errorMessages,
			})
		} else {
			
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
				"error":   err.Error(),
			})
		}
	
	}
	h.jawabanPesertaService.UpdateSkor(jawaban_pesertaRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "jawaban_peserta updated",
	})
}

func(h *jwabanPesertaHandler) GetJawabanByUserAndQuiz(ctx *gin.Context) {
    idUser := ctx.Query("id_user")
	idQuiz := ctx.Query("id_quiz")

    jawaban_peserta, err := h.jawabanPesertaService.FindByIDUserAndIDQuiz(idUser, idQuiz)
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

    var jawaban_pesertaResponse []response.JawabanPesertaAdminResponse

    jawaban_pesertaResponseItem := response.JawabanPesertaAdminResponse{
		ID:         jawaban_peserta.ID,
		IdUser:     jawaban_peserta.IdUser,
		IdQuiz:     jawaban_peserta.IdQuiz,
		IdPertanyaan: jawaban_peserta.IdPertanyaan,
		Skor: 	 jawaban_peserta.Skor,
	
	}
	jawaban_pesertaResponse = append(jawaban_pesertaResponse, jawaban_pesertaResponseItem)

    ctx.JSON(http.StatusOK, gin.H{
        "message": "success",
        "data":    jawaban_pesertaResponse,
    })
}

func (h *jwabanPesertaHandler) GetJawabanPesertas(ctx *gin.Context) {
	jawaban_peserta, err := h.jawabanPesertaService.FindAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	var jawaban_pesertaResponse []response.JawabanPesertaAdminResponse

	for _, p := range jawaban_peserta {
		
		total_soal,_ := h.jawabanPesertaService.CountQuestion(p.IdQuiz)
		jumlah_benar := h.jawabanPesertaService.CountRightAnswer(p.IdQuiz, p.IdUser, p.IdPertanyaan, p.JawabanPeserta)
		jawaban_pesertaResponse = append(jawaban_pesertaResponse, response.NewJawabanPesertaAdminResponse(p, total_soal, jumlah_benar))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    jawaban_pesertaResponse,
	})
}

func (h *jwabanPesertaHandler) GetJawabanPeserta(ctx *gin.Context) {
	idString := ctx.Param("id")
	idUint, _ := strconv.ParseUint(idString, 10, 64)
	id := uint(idUint)

	jawaban_peserta, err := h.jawabanPesertaService.FindByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	jawaban_pesertaResponse := response.JawabanPesertaResponse{
		ID:         jawaban_peserta.ID,
		IdUser:     jawaban_peserta.IdUser,
		IdQuiz:     jawaban_peserta.IdQuiz,
		IdPertanyaan: jawaban_peserta.IdPertanyaan,
		JawabanPeserta: jawaban_peserta.JawabanPeserta,
		Skor:       jawaban_peserta.Skor,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    jawaban_pesertaResponse,
	})
}

func (h *jwabanPesertaHandler) CreateJawabanPeserta(ctx *gin.Context) {
	var jawaban_pesertaRequest request.JawabanPesertaRequest
	err := ctx.ShouldBindJSON(&jawaban_pesertaRequest)
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

	jawaban_peserta, err := h.jawabanPesertaService.Create(jawaban_pesertaRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jawaban_pesertaResponse := response.JawabanPesertaResponse{
		ID:         jawaban_peserta.ID,
		IdUser:     jawaban_peserta.IdUser,
		IdQuiz:     jawaban_peserta.IdQuiz,
		IdPertanyaan: jawaban_peserta.IdPertanyaan,
		JawabanPeserta: jawaban_peserta.JawabanPeserta,
		Skor:       jawaban_peserta.Skor,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    jawaban_pesertaResponse,
	})
}

func (h *jwabanPesertaHandler) UpdateJawabanPeserta(ctx *gin.Context) {
	var jawaban_pesertaRequest request.JawabanPesertaRequest
	err := ctx.ShouldBindJSON(&jawaban_pesertaRequest)
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

	h.jawabanPesertaService.Update(id, jawaban_pesertaRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "jawaban_peserta updated",
	})
}

func (h *jwabanPesertaHandler) DeleteJawabanPeserta(ctx *gin.Context) {
	idString := ctx.Param("id")
	idUint, _ := strconv.ParseUint(idString, 10, 64)
	id := uint(idUint)

	h.jawabanPesertaService.Delete(id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "jawaban_peserta deleted",
	})
}