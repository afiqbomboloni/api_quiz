package entity

import (

	"gorm.io/gorm"
)

type Pertanyaan struct {
	gorm.Model
    ID            uint `gorm:"primary_key"`
    Pertanyaan    string `gorm:"type:longtext;" json:"pertanyaan"`
    OpsiJawaban   string `gorm:"type:longtext;" json:"opsi_jawaban"`
    JawabanBenar  int `gorm:"not null;" json:"jawaban_benar"`
    IdQuiz        uint `gorm:"not null;" json:"id_quiz"`
    Quiz          Quiz `gorm:"foreignKey:IdQuiz"`
}