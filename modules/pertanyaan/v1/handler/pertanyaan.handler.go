package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/afiqbomboloni/api_quiz/modules/pertanyaan/v1/service"
	"github.com/afiqbomboloni/api_quiz/request"
	"github.com/afiqbomboloni/api_quiz/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type pertanyaanHandler struct {
	pertanyaanService service.PertanyaanService
}

func NewPertanyaanHandler(pertanyaanService service.PertanyaanService) *pertanyaanHandler {
	return &pertanyaanHandler{pertanyaanService}
}

func (h *pertanyaanHandler) GetPertanyaans(ctx *gin.Context) {
	pertanyaans, err := h.pertanyaanService.FindAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	var pertanyaansResponse []response.PertanyaanResponse
	for _, p := range pertanyaans {
		pertanyaansResponse = append(pertanyaansResponse, response.PertanyaanResponse{
			ID:           p.ID,
			Pertanyaan:   p.Pertanyaan,
			OpsiJawaban:  p.OpsiJawaban,
			JawabanBenar: p.JawabanBenar,
			IdQuiz:       p.IdQuiz,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    pertanyaansResponse,
	})
}

func (h *pertanyaanHandler) GetPertanyaan(ctx *gin.Context) {
	idString := ctx.Param("id")
	idUint, _ := strconv.ParseUint(idString, 10, 64)
	id := uint(idUint)

	pertanyaan, err := h.pertanyaanService.FindByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	pertanyaanResponse := response.PertanyaanResponse{
		ID: 		 pertanyaan.ID,
		Pertanyaan:  pertanyaan.Pertanyaan,
		OpsiJawaban: pertanyaan.OpsiJawaban,
		JawabanBenar: pertanyaan.JawabanBenar,
		IdQuiz:      pertanyaan.IdQuiz,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    pertanyaanResponse,
	})
}

func (h *pertanyaanHandler) GetPertanyaanByQuiz(ctx *gin.Context) {
	idString := ctx.Param("id")
	idUint, _ := strconv.ParseUint(idString, 10, 64)
	id := uint(idUint)

	pertanyaans, err := h.pertanyaanService.FindByQuiz(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	var pertanyaansResponse []response.PertanyaanResponse
	for _, p := range pertanyaans {
		pertanyaansResponse = append(pertanyaansResponse, response.PertanyaanResponse{
			ID:           p.ID,
			Pertanyaan:   p.Pertanyaan,
			OpsiJawaban:  p.OpsiJawaban,
			JawabanBenar: p.JawabanBenar,
			IdQuiz:       p.IdQuiz,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    pertanyaansResponse,
	})
}

func (h *pertanyaanHandler) CreatePertanyaan(ctx *gin.Context) {
	var pertanyaanRequest request.PertanyaanRequest
	err := ctx.ShouldBindJSON(&pertanyaanRequest)
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

	h.pertanyaanService.Create(pertanyaanRequest)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "new pertanyaan created",
	})
}

func (h *pertanyaanHandler) UpdatePertanyaan(ctx *gin.Context) {
	var pertanyaanRequest request.PertanyaanRequest
	err := ctx.ShouldBindJSON(&pertanyaanRequest)
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

	h.pertanyaanService.Update(id, pertanyaanRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pertanyaan updated",
	},
)
}

func (h *pertanyaanHandler) DeletePertanyaan(ctx *gin.Context) {
	idString := ctx.Param("id")
	idUint, _ := strconv.ParseUint(idString, 10, 64)
	id := uint(idUint)

	h.pertanyaanService.Delete(id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pertanyaan deleted",
	})
}