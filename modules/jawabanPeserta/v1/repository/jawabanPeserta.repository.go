package repository

import (
	// "fmt"

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
	CountQuestion(IdQuiz uint) (int64, error)
	CountRightAnswer(IdQuiz uint, IdUser uint, IdPertanyaan uint, JawabanPeserta int) (int64)
	FindAllWIthUniqueUser() ([]entity.JawabanPeserta, error)
	UpdateSkor(jawabanPeserta entity.JawabanPeserta) (entity.JawabanPeserta, error)
	FindByIDUserAndIDQuiz(IDUser, IDQuiz string) (entity.JawabanPeserta, error)

}

type jawabanPesertaRepository struct {
	db *gorm.DB
}

func NewJawabanPesertaRepository(db *gorm.DB) *jawabanPesertaRepository {
	return &jawabanPesertaRepository{db}
}

func (r *jawabanPesertaRepository) FindAll() ([]entity.JawabanPeserta, error) {
	var jawabanPesertas []entity.JawabanPeserta

	err := r.db.Model(&entity.JawabanPeserta{}).Debug().Preload("User").Preload("Quiz").Find(&jawabanPesertas).Error
	if err != nil {
		return nil, err
	}
	

	return jawabanPesertas, nil
}

func (r *jawabanPesertaRepository) UpdateSkor(jawabanPeserta entity.JawabanPeserta) (entity.JawabanPeserta, error) {
	err := r.db.Debug().Model(&entity.JawabanPeserta{}).Where("id_user = ? AND id_quiz = ?", jawabanPeserta.IdUser, jawabanPeserta.IdQuiz).Updates(&entity.JawabanPeserta{Skor: jawabanPeserta.Skor}).Error
	if err != nil {
		return entity.JawabanPeserta{}, err
	}
	return jawabanPeserta, nil

}

func (r *jawabanPesertaRepository) FindAllWIthUniqueUser() ([]entity.JawabanPeserta, error) {
	var jawabanPesertas []entity.JawabanPeserta
	subQuery := r.db.Table("jawaban_peserta").Select("MIN(id) as id").Group("id_user, id_quiz")
	err := r.db.Table("jawaban_peserta").Debug().Preload("User").Preload("Quiz").Joins("JOIN (?) AS t ON jawaban_peserta.id = t.id", subQuery).Find(&jawabanPesertas).Error
	if err != nil {
		return nil, err
	}
	return jawabanPesertas, nil
}



func (r *jawabanPesertaRepository) CountQuestion(IdQuiz uint) (int64, error) {
	var count int64
	err := r.db.Model(&entity.JawabanPeserta{}).Where("id_quiz = ?", IdQuiz).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil

}

func (r *jawabanPesertaRepository) FindByIDUserAndIDQuiz(IDUser, IDQuiz string) (entity.JawabanPeserta, error) {
	var jawabanPeserta entity.JawabanPeserta
	
	err := r.db.Where("id_user = ? AND id_quiz = ?", IDUser, IDQuiz).First(&jawabanPeserta).Error
	if err != nil {
		return entity.JawabanPeserta{}, err
	}
	return jawabanPeserta, nil
}

func (r *jawabanPesertaRepository) CountRightAnswer(IdQuiz uint, IdUser uint, IdPertanyaan uint, JawabanPeserta int) (int64) {
	var count int64
	r.db.Model(&entity.JawabanPeserta{}).Where("id_quiz = ? AND id_user = ? AND id_pertanyaan = ? AND jawaban_peserta = ?", IdQuiz, IdUser, IdPertanyaan, JawabanPeserta).Count(&count)
	return count

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