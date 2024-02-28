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

func (h *jwabanPesertaHandler) GetJawabanPesertas(ctx *gin.Context) {
	jawaban_peserta, err := h.jawabanPesertaService.FindAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	var jawaban_pesertaResponse []response.JawabanPesertaResponse
	for _, p := range jawaban_peserta {
		jawaban_pesertaResponse = append(jawaban_pesertaResponse, response.JawabanPesertaResponse{
			ID:         p.ID,
			IdUser:     p.IdUser,
			IdQuiz:     p.IdQuiz,
			IdPertanyaan: p.IdPertanyaan,
			JawabanPeserta: p.JawabanPeserta,
			Skor:       p.Skor,
		})
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