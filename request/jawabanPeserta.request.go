package request

/*
type JawabanPeserta struct {
	gorm.Model
    ID            uint `gorm:"primary_key"`
    IdUser        uint `gorm:"not null;" json:"id_user"`
    IdQuiz        uint `gorm:"not null;" json:"id_quiz"`
    IdPertanyaan  uint `gorm:"not null;" json:"id_pertanyaan"`
    JawabanPeserta int `gorm:"not null;" json:"jawaban_peserta"`
    Skor          int `json:"skor"`
    User          User `gorm:"foreignKey:IdUser"`
}
*/
type JawabanPesertaRequest struct {
	ID            uint `json:"id"`
	IdUser        uint `json:"id_user"`
	IdQuiz        uint `json:"id_quiz"`
	IdPertanyaan  uint `json:"id_pertanyaan"`
	JawabanPeserta int `json:"jawaban_peserta"`
	Skor          int `json:"skor"`
}

type UpdateSkorRequest struct {
	IdUser        uint `json:"id_user"`
	IdQuiz        uint `json:"id_quiz"`
	Skor          int `json:"skor"`
}