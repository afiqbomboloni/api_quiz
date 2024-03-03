package repository

import (
	"time"

	"github.com/afiqbomboloni/api_quiz/entity"
	"gorm.io/gorm"
)

type QuizRepository interface {
	FindAll(limit, page uint64, isNotExpired string) ([]entity.Quiz, error)
	FindByID(id uint) (entity.Quiz, error)
	Create(quiz entity.Quiz) (entity.Quiz, error)
	Update(quiz entity.Quiz) (entity.Quiz, error)
	Delete(quiz entity.Quiz) (entity.Quiz, error)
	CountQuestion(IdQuiz uint) (int64, error)
}

type quizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *quizRepository {
	return &quizRepository{db}
}

func(r *quizRepository) CountQuestion(idQuiz uint) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Pertanyaan{}).Where("id_quiz = ?", idQuiz).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *quizRepository) FindAll(limit,page uint64, isNotExpired string) ([]entity.Quiz, error) {
	var quizzes []entity.Quiz

	
	if(isNotExpired == "true") {
		err := r.db.Limit(int(limit)).Offset(int((page - 1) * limit)).Debug().Preload("Pertanyaan").Where("waktu_mulai < ? AND waktu_selesai > ?", time.Now(), time.Now()).Find(&quizzes).Error
		if err != nil {
			return nil, err
		}
		return quizzes, nil
	
	} 

	err := r.db.Limit(int(limit)).Offset(int((page - 1) * limit)).Debug().Preload("Pertanyaan").Find(&quizzes).Error
	if err != nil {
		return nil, err
	}
	
	return quizzes, nil
	
	
	
}

func (r *quizRepository) FindByID(id uint) (entity.Quiz, error) {
	var quiz entity.Quiz

	err := r.db.First(&quiz, id).Error
	if err != nil {
		return entity.Quiz{}, err
	}

	return quiz, nil
}

func (r *quizRepository) Create(quiz entity.Quiz) (entity.Quiz, error) {
	err := r.db.Debug().Create(&quiz).Error
	if err != nil {
		return entity.Quiz{}, err
	}

	return quiz, nil
}

func (r *quizRepository) Update(quiz entity.Quiz) (entity.Quiz, error) {
	err := r.db.Debug().Updates(&quiz).Error
	if err != nil {
		return entity.Quiz{}, err
	}

	return quiz, nil
}

func (r *quizRepository) Delete(quiz entity.Quiz) (entity.Quiz, error) {
	err := r.db.Debug().Delete(&quiz).Error
	if err != nil {
		return entity.Quiz{}, err
	}

	return quiz, nil
}