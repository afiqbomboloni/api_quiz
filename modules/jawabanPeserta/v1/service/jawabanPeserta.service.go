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
	CountQuestion(IdQuiz uint) (int64, error)
	CountRightAnswer(IdQuiz uint, IdUser uint, IdPertanyaan uint, JawabanPeserta int) (int64)
	FindAllWIthUniqueUser() ([]entity.JawabanPeserta, error)
	UpdateSkor(jawabanPesertaRequest request.UpdateSkorRequest) (entity.JawabanPeserta, error)
	FindByIDUserAndIDQuiz(IDUser, IDQuiz string) (entity.JawabanPeserta, error)

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

func (s *jawabanPesertaService) CountQuestion(IdQuiz uint) (int64, error) {
	return s.jawabanPesertaRepository.CountQuestion(IdQuiz)

}

func (s *jawabanPesertaService) CountRightAnswer(IdQuiz uint, IdUser uint, IdPertanyaan uint, JawabanPeserta int) (int64) {
	return s.jawabanPesertaRepository.CountRightAnswer(IdQuiz, IdUser, IdPertanyaan, JawabanPeserta)


}

func (s *jawabanPesertaService) FindAllWIthUniqueUser() ([]entity.JawabanPeserta, error) {
	return s.jawabanPesertaRepository.FindAllWIthUniqueUser()

}

func (s *jawabanPesertaService) UpdateSkor(jawabanPesertaRequest request.UpdateSkorRequest) (entity.JawabanPeserta, error) {
    jawabanPeserta := entity.JawabanPeserta{
        IdUser: jawabanPesertaRequest.IdUser,
        IdQuiz: jawabanPesertaRequest.IdQuiz,
        Skor:   jawabanPesertaRequest.Skor,
       
    }
    return s.jawabanPesertaRepository.UpdateSkor(jawabanPeserta)
}

func (s *jawabanPesertaService) FindByIDUserAndIDQuiz(IDUser, IDQuiz string) (entity.JawabanPeserta, error) {
	return s.jawabanPesertaRepository.FindByIDUserAndIDQuiz(IDUser, IDQuiz)

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