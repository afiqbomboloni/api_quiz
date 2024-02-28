package entity

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
    ID           uint `gorm:"primary_key"`
    Judul        string `gorm:"size:255;not null;" json:"judul"`
    Deskripsi    string `gorm:"type:longtext;" json:"deskripsi"`
    WaktuMulai   time.Time `gorm:"not null;" json:"waktu_mulai"`
    WaktuSelesai time.Time `gorm:"not null;" json:"waktu_selesai"`
    Pertanyaan   []Pertanyaan `gorm:"foreignKey:IdQuiz"`
}