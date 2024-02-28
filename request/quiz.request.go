package request

type QuizRequest struct {
	Judul		string `json:"judul"`
	Deskripsi	string `json:"deskripsi"`
	WaktuMulai	string `json:"waktu_mulai"`
	WaktuSelesai	string `json:"waktu_selesai"`
}
