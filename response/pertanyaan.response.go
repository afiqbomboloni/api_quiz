package response

type PertanyaanResponse struct {
	ID           uint   `json:"id"`
	Pertanyaan   string `json:"pertanyaan"`
	OpsiJawaban  string `json:"opsi_jawaban"`
	JawabanBenar int `json:"jawaban_benar"`
	IdQuiz       uint   `json:"id_quiz"`
}