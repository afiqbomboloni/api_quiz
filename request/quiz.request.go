package request

type QuizRequest struct {
	Judul		string `json:"judul"`
	Deskripsi	string `json:"deskripsi"`
	WaktuMulai	string `json:"waktu_mulai"`
	WaktuSelesai	string `json:"waktu_selesai"`
}

/*
example request

{
	"judul": "Quiz 1",
	"deskripsi": "Quiz 1",
	"waktu_mulai": "2021-01-01T00:00:00Z",
	"waktu_selesai": "2023-01-01T00:00:00Z"
}

*/
