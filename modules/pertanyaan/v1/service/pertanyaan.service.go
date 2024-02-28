package service

import (
	"github.com/afiqbomboloni/api_quiz/entity"
	"github.com/afiqbomboloni/api_quiz/modules/pertanyaan/v1/repository"
	"github.com/afiqbomboloni/api_quiz/request"
)

type PertanyaanService interface {
	FindAll() ([]entity.Pertanyaan, error)
	FindByID(id uint) (entity.Pertanyaan, error)
	Create(pertanyaanRequest request.PertanyaanRequest) (entity.Pertanyaan, error)
	Update(id uint, pertanyaanRequest request.PertanyaanRequest) (entity.Pertanyaan, error)
	Delete(id uint) (entity.Pertanyaan, error)
	FindByQuiz(id uint) ([]entity.Pertanyaan, error)
}

type pertanyaanService struct {
	pertanyaanRepository repository.PertanyaanRepository
}

func NewPertanyaanService(pertanyaanRepository repository.PertanyaanRepository) *pertanyaanService {
	return &pertanyaanService{pertanyaanRepository}
}

func (s *pertanyaanService) FindAll() ([]entity.Pertanyaan, error) {
	pertanyaans, err := s.pertanyaanRepository.FindAll()
	return pertanyaans, err
}

func (s *pertanyaanService) FindByQuiz(id uint) ([]entity.Pertanyaan, error) {
	pertanyaans, err := s.pertanyaanRepository.FindByQuiz(id)
	return pertanyaans, err
}

func (s *pertanyaanService) FindByID(id uint) (entity.Pertanyaan, error) {
	pertanyaan, err := s.pertanyaanRepository.FindByID(id)

	return pertanyaan, err
}

func (s *pertanyaanService) Create(pertanyaanRequest request.PertanyaanRequest) (entity.Pertanyaan, error) {
	pertanyaan := entity.Pertanyaan{
		Pertanyaan:   pertanyaanRequest.Pertanyaan,
		OpsiJawaban:  pertanyaanRequest.OpsiJawaban,
		JawabanBenar: pertanyaanRequest.JawabanBenar,
		IdQuiz:       pertanyaanRequest.IdQuiz,
	}
	newPertanyaan, err := s.pertanyaanRepository.Create(pertanyaan)
	return newPertanyaan, err
}

func (s *pertanyaanService) Update(id uint, pertanyaanRequest request.PertanyaanRequest) (entity.Pertanyaan, error) {
	pertanyaan, err := s.pertanyaanRepository.FindByID(id)
	if err != nil {
		return entity.Pertanyaan{}, err
	}

	pertanyaan.Pertanyaan = pertanyaanRequest.Pertanyaan
	pertanyaan.OpsiJawaban = pertanyaanRequest.OpsiJawaban
	pertanyaan.JawabanBenar = pertanyaanRequest.JawabanBenar
	pertanyaan.IdQuiz = pertanyaanRequest.IdQuiz

	updatedPertanyaan, err := s.pertanyaanRepository.Update(pertanyaan)
	return updatedPertanyaan, err
}

func (s *pertanyaanService) Delete(id uint) (entity.Pertanyaan, error) {
	pertanyaan, err := s.pertanyaanRepository.FindByID(id)
	if err != nil {
		return entity.Pertanyaan{}, err
	}

	deletedPertanyaan, err := s.pertanyaanRepository.Delete(pertanyaan)
	return deletedPertanyaan, err
}