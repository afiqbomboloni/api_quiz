package repository

import (
	"github.com/afiqbomboloni/api_quiz/entity"
	"gorm.io/gorm"
)

type JawabanPesertaRepository interface {
	FindAll() ([]entity.JawabanPeserta, error)
	FindByID(ID uint) (entity.JawabanPeserta, error)
	FindByIDPeserta(ID uint) ([]entity.JawabanPeserta, error)
	FindByIDSoal(ID uint) ([]entity.JawabanPeserta, error)
	FindByIDPesertaAndIDSoal(IDPeserta uint, IDSoal uint) (entity.JawabanPeserta, error)
	Create(jawabanPeserta entity.JawabanPeserta) (entity.JawabanPeserta, error)
	Update(jawabanPeserta entity.JawabanPeserta) (entity.JawabanPeserta, error)
	Delete(jawabanPeserta entity.JawabanPeserta) (entity.JawabanPeserta, error)
}

type jawabanPesertaRepository struct {
	db *gorm.DB
}

func NewJawabanPesertaRepository(db *gorm.DB) *jawabanPesertaRepository {
	return &jawabanPesertaRepository{db}
}

func (r *jawabanPesertaRepository) FindAll() ([]entity.JawabanPeserta, error) {
	var jawabanPesertas []entity.JawabanPeserta

	err := r.db.Find(&jawabanPesertas).Error
	if err != nil {
		return nil, err
	}

	return jawabanPesertas, nil
}

func (r *jawabanPesertaRepository) FindByID(ID uint) (entity.JawabanPeserta, error) {
	var jawabanPeserta entity.JawabanPeserta
	err := r.db.First(&jawabanPeserta, ID).Error
	if err != nil {
		return entity.JawabanPeserta{}, err
	}
	return jawabanPeserta, nil
}

func (r *jawabanPesertaRepository) FindByIDPeserta(ID uint) ([]entity.JawabanPeserta, error) {
	var jawabanPesertas []entity.JawabanPeserta
	err := r.db.Where("id_peserta = ?", ID).Find(&jawabanPesertas).Error
	if err != nil {
		return nil, err
	}
	return jawabanPesertas, nil
}

func (r *jawabanPesertaRepository) FindByIDSoal(ID uint) ([]entity.JawabanPeserta, error) {
	var jawabanPesertas []entity.JawabanPeserta
	err := r.db.Where("id_soal = ?", ID).Find(&jawabanPesertas).Error
	if err != nil {
		return nil, err
	}
	return jawabanPesertas, nil
}

func (r *jawabanPesertaRepository) FindByIDPesertaAndIDSoal(IDPeserta uint, IDSoal uint) (entity.JawabanPeserta, error) {
	var jawabanPeserta entity.JawabanPeserta
	err := r.db.Where("id_peserta = ? AND id_soal = ?", IDPeserta, IDSoal).First(&jawabanPeserta).Error
	if err != nil {
		return entity.JawabanPeserta{}, err
	}
	return jawabanPeserta, nil
}

func (r *jawabanPesertaRepository) Create(jawabanPeserta entity.JawabanPeserta) (entity.JawabanPeserta, error) {
	err := r.db.Debug().Create(&jawabanPeserta).Error
	if err != nil {
		return entity.JawabanPeserta{}, err
	}
	return jawabanPeserta, nil
}

func (r *jawabanPesertaRepository) Update(jawabanPeserta entity.JawabanPeserta) (entity.JawabanPeserta, error) {
	err := r.db.Debug().Updates(&jawabanPeserta).Error
	if err != nil {
		return entity.JawabanPeserta{}, err
	}
	return jawabanPeserta, nil
}

func (r *jawabanPesertaRepository) Delete(jawabanPeserta entity.JawabanPeserta) (entity.JawabanPeserta, error) {
	err := r.db.Debug().Delete(&jawabanPeserta).Error
	if err != nil {
		return entity.JawabanPeserta{}, err
	}
	return jawabanPeserta, nil
}