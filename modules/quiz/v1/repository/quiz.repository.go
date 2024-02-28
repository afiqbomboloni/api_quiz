package repository

import (
	"github.com/afiqbomboloni/api_quiz/entity"
	"gorm.io/gorm"
)

type QuizRepository interface {
	FindAll() ([]entity.Quiz, error)
	FindByID(id uint) (entity.Quiz, error)
	Create(quiz entity.Quiz) (entity.Quiz, error)
	Update(quiz entity.Quiz) (entity.Quiz, error)
	Delete(quiz entity.Quiz) (entity.Quiz, error)
}

type quizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *quizRepository {
	return &quizRepository{db}
}

func (r *quizRepository) FindAll() ([]entity.Quiz, error) {
	var quizzes []entity.Quiz

	err := r.db.Find(&quizzes).Error
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