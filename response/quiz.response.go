package response

type QuizResponse struct {
	ID 		uint   `json:"id"`
	Judul		string `json:"judul"`
	Deskripsi	string `json:"deskripsi"`
	WaktuMulai	string `json:"waktu_mulai"`
	WaktuSelesai	string `json:"waktu_selesai"`
}