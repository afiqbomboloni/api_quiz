package service

import (
	"time"

	"github.com/afiqbomboloni/api_quiz/entity"
	"github.com/afiqbomboloni/api_quiz/modules/quiz/v1/repository"
	"github.com/afiqbomboloni/api_quiz/request"
)

type QuizService interface {
	FindAll() ([]entity.Quiz, error)
	FindByID(id uint) (entity.Quiz, error)
	Create(quizRequest request.QuizRequest) (entity.Quiz, error)
	Update(id uint, quizRequest request.QuizRequest) (entity.Quiz, error)
	Delete(id uint) (entity.Quiz, error)
}

type quizService struct {
	quizRepository repository.QuizRepository
}

func NewQuizService(quizRepository repository.QuizRepository) *quizService {
	return &quizService{quizRepository}
}

func (s *quizService) FindAll() ([]entity.Quiz, error) {
	quizzes, err := s.quizRepository.FindAll()
	return quizzes, err
}

func (s *quizService) FindByID(id uint) (entity.Quiz, error) {
	quiz, err := s.quizRepository.FindByID(id)
	return quiz, err
}

func (s *quizService) Create(quizRequest request.QuizRequest) (entity.Quiz, error) {
	layout := "2006-01-02T15:04:05Z"
	waktuMulai, err := time.Parse(layout, quizRequest.WaktuMulai)
	if err != nil {
		return entity.Quiz{}, err
	
	}
	waktuSelesai, err := time.Parse(layout, quizRequest.WaktuSelesai)
	if err != nil {
		return entity.Quiz{}, err
	
	}
	quiz := entity.Quiz{
		Judul:       quizRequest.Judul,
		Deskripsi:   quizRequest.Deskripsi,
		WaktuMulai: waktuMulai,
		WaktuSelesai: waktuSelesai,
	}
	newQuiz, err := s.quizRepository.Create(quiz)
	return newQuiz, err
}

func (s *quizService) Update(id uint, quizRequest request.QuizRequest) (entity.Quiz, error) {
	quiz, err := s.quizRepository.FindByID(id)

	if err != nil {
		return entity.Quiz{}, err
	}
	
	quiz.Judul = quizRequest.Judul
	quiz.Deskripsi = quizRequest.Deskripsi
	layout := "2006-01-02T15:04:05Z"
	waktuMulai, err := time.Parse(layout, quizRequest.WaktuMulai)
	if err != nil {
		return entity.Quiz{}, err
	
	}
	quiz.WaktuMulai = waktuMulai
	waktuSelesai, err := time.Parse(layout, quizRequest.WaktuSelesai)

	if err != nil {
		return entity.Quiz{}, err
	
	}
	quiz.WaktuSelesai = waktuSelesai

	updatedQuiz, err := s.quizRepository.Update(quiz)

	return updatedQuiz, err
}

func (s *quizService) Delete(id uint) (entity.Quiz, error) {
	quiz, err := s.quizRepository.FindByID(id)

	if err != nil {
		return entity.Quiz{}, err
	}

	deletedQuiz, err := s.quizRepository.Delete(quiz)

	return deletedQuiz, err
}