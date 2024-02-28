package service

import (
	"github.com/afiqbomboloni/api_quiz/entity"
	"github.com/afiqbomboloni/api_quiz/modules/jawabanPeserta/v1/repository"
	"github.com/afiqbomboloni/api_quiz/request"
)

type JawabanPesertaService interface {
	FindAll() ([]entity.JawabanPeserta, error)
	FindByID(ID uint) (entity.JawabanPeserta, error)
	FindByIDPeserta(IDPeserta uint) ([]entity.JawabanPeserta, error)
	FindByIDSoal(IDSoal uint) ([]entity.JawabanPeserta, error)
	FindByIDPesertaAndIDSoal(IDPeserta uint, IDSoal uint) (entity.JawabanPeserta, error)
	Create(jawabanPesertaRequest request.JawabanPesertaRequest) (entity.JawabanPeserta, error)
	Update(ID uint, jawabanPesertaRequest request.JawabanPesertaRequest) (entity.JawabanPeserta, error)
	Delete(ID uint) (entity.JawabanPeserta, error)
}

type jawabanPesertaService struct {
	jawabanPesertaRepository repository.JawabanPesertaRepository
}

func NewJawabanPesertaService(jawabanPesertaRepository repository.JawabanPesertaRepository) *jawabanPesertaService {
	return &jawabanPesertaService{jawabanPesertaRepository}
}

func (s *jawabanPesertaService) FindAll() ([]entity.JawabanPeserta, error) {
	return s.jawabanPesertaRepository.FindAll()
}

func (s *jawabanPesertaService) FindByID(ID uint) (entity.JawabanPeserta, error) {
	jawaban_peserta, err := s.jawabanPesertaRepository.FindByID(ID)
	if err != nil {
		return entity.JawabanPeserta{}, err
	}

	return jawaban_peserta, nil
}

func (s *jawabanPesertaService) FindByIDPeserta(IDPeserta uint) ([]entity.JawabanPeserta, error) {
	jawaban_pesertas, err := s.jawabanPesertaRepository.FindByIDPeserta(IDPeserta)
	if err != nil {
		return nil, err
	}

	return jawaban_pesertas, nil
}

func (s *jawabanPesertaService) FindByIDSoal(IDSoal uint) ([]entity.JawabanPeserta, error) {
	jawaban_pesertas, err := s.jawabanPesertaRepository.FindByIDSoal(IDSoal)
	if err != nil {
		return nil, err
	}

	return jawaban_pesertas, nil
}

func (s *jawabanPesertaService) FindByIDPesertaAndIDSoal(IDPeserta uint, IDSoal uint) (entity.JawabanPeserta, error) {
	jawaban_peserta, err := s.jawabanPesertaRepository.FindByIDPesertaAndIDSoal(IDPeserta, IDSoal)
	if err != nil {
		return entity.JawabanPeserta{}, err
	}

	return jawaban_peserta, nil
}

func (s *jawabanPesertaService) Create(jawabanPesertaRequest request.JawabanPesertaRequest) (entity.JawabanPeserta, error) {
	jawaban_peserta := entity.JawabanPeserta{
		IdUser:        jawabanPesertaRequest.IdUser,
		IdQuiz:        jawabanPesertaRequest.IdQuiz,
		IdPertanyaan:  jawabanPesertaRequest.IdPertanyaan,
		JawabanPeserta: jawabanPesertaRequest.JawabanPeserta,
		Skor:          jawabanPesertaRequest.Skor,
	}

	newJawabanPeserta, err := s.jawabanPesertaRepository.Create(jawaban_peserta)
	return newJawabanPeserta, err
}

func (s *jawabanPesertaService) Update(ID uint, jawabanPesertaRequest request.JawabanPesertaRequest) (entity.JawabanPeserta, error) {
	jawaban_peserta, err := s.jawabanPesertaRepository.FindByID(ID)
	if err != nil {
		return entity.JawabanPeserta{}, err
	}

	jawaban_peserta.IdUser = jawabanPesertaRequest.IdUser
	jawaban_peserta.IdQuiz = jawabanPesertaRequest.IdQuiz
	jawaban_peserta.IdPertanyaan = jawabanPesertaRequest.IdPertanyaan
	jawaban_peserta.JawabanPeserta = jawabanPesertaRequest.JawabanPeserta
	jawaban_peserta.Skor = jawabanPesertaRequest.Skor
	

	updatedJawabanPeserta, err := s.jawabanPesertaRepository.Update(jawaban_peserta)
	return updatedJawabanPeserta, err
}

func (s *jawabanPesertaService) Delete(ID uint) (entity.JawabanPeserta, error) {
	jawaban_peserta, err := s.jawabanPesertaRepository.FindByID(ID)
	if err != nil {
		return entity.JawabanPeserta{}, err
	}

	deletedJawabanPeserta, err := s.jawabanPesertaRepository.Delete(jawaban_peserta)
	return deletedJawabanPeserta, err
}