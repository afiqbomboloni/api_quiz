package repository

import (
	"github.com/afiqbomboloni/api_quiz/entity"
	"gorm.io/gorm"
)

type PertanyaanRepository interface {
	FindAll() ([]entity.Pertanyaan, error)
	FindByID(id uint) (entity.Pertanyaan, error)
	Create(pertanyaan entity.Pertanyaan) (entity.Pertanyaan, error)
	Update(pertanyaan entity.Pertanyaan) (entity.Pertanyaan, error)
	Delete(pertanyaan entity.Pertanyaan) (entity.Pertanyaan, error)
	FindByQuiz(id uint) ([]entity.Pertanyaan, error)
}

type pertanyaanRepository struct {
	db *gorm.DB
}

func NewPertanyaanRepository(db *gorm.DB) *pertanyaanRepository {
	return &pertanyaanRepository{db}
}

func (r *pertanyaanRepository) FindAll() ([]entity.Pertanyaan, error) {
	var pertanyaans []entity.Pertanyaan

	err := r.db.Find(&pertanyaans).Error
	if err != nil {
		return nil, err
	}

	return pertanyaans, nil
}

func (r *pertanyaanRepository) FindByQuiz(id uint) ([]entity.Pertanyaan, error) {
	var pertanyaans []entity.Pertanyaan

	err := r.db.Where("id_quiz = ?", id).Find(&pertanyaans).Error
	if err != nil {
		return nil, err
	}

	return pertanyaans, nil
}

func (r *pertanyaanRepository) FindByID(id uint) (entity.Pertanyaan, error) {
	var pertanyaan entity.Pertanyaan
	err := r.db.First(&pertanyaan, id).Error
	if err != nil {
		return entity.Pertanyaan{}, err
	}
	return pertanyaan, nil
}

func (r *pertanyaanRepository) Create(pertanyaan entity.Pertanyaan) (entity.Pertanyaan, error) {
	err := r.db.Debug().Create(&pertanyaan).Error
	if err != nil {
		return entity.Pertanyaan{}, err
	}
	return pertanyaan, nil
}

func (r *pertanyaanRepository) Update(pertanyaan entity.Pertanyaan) (entity.Pertanyaan, error) {
	err := r.db.Debug().Updates(&pertanyaan).Error
	if err != nil {
		return entity.Pertanyaan{}, err
	}
	return pertanyaan, nil
}

func (r *pertanyaanRepository) Delete(pertanyaan entity.Pertanyaan) (entity.Pertanyaan, error) {
	err := r.db.Debug().Delete(&pertanyaan).Error
	if err != nil {
		return entity.Pertanyaan{}, err
	}
	return pertanyaan, nil
}
