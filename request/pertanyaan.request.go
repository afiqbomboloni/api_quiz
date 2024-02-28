package request

type PertanyaanRequest struct {
	Pertanyaan string `json:"pertanyaan" binding:"required"`
	OpsiJawaban string `json:"opsi_jawaban" binding:"required"`
	JawabanBenar int `json:"jawaban_benar" binding:"required"`
	IdQuiz uint `json:"id_quiz" binding:"required"`
}