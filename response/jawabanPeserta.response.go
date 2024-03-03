package response

import "github.com/afiqbomboloni/api_quiz/entity"

type JawabanPesertaResponse struct {
	ID             uint `json:"id"`
	IdUser         uint `json:"id_user"`
	IdQuiz         uint `json:"id_quiz"`
	IdPertanyaan   uint `json:"id_pertanyaan"`
	JawabanPeserta int  `json:"jawaban_peserta"`
	Skor           int  `json:"skor"`
}

type JawabanPesertaAdminResponse struct {
	ID     uint `json:"id"`
	IdUser uint `json:"id_user"`
	User   User `json:"user"`
	IdQuiz uint `json:"id_quiz"`
	Quiz   Quiz `json:"quiz"`
	IdPertanyaan   uint  `json:"id_pertanyaan"`
	JawabanPeserta int   `json:"jawaban_peserta"`
	JumlahBenar    int64 `json:"jumlah_benar"`
	TotalSoal      int64 `json:"total_soal"`
	Skor           int   `json:"skor"`
}

type User struct {
    Nama string `json:"nama"`
}

type Quiz struct {
    Judul string `json:"judul"`
}

func NewJawabanPesertaAdminResponse(jawaban_peserta entity.JawabanPeserta, total_soal, jumlah_benar int64) JawabanPesertaAdminResponse {
	res := JawabanPesertaAdminResponse{
		ID:             jawaban_peserta.ID,
		IdUser:         jawaban_peserta.IdUser,
		IdQuiz:         jawaban_peserta.IdQuiz,
		IdPertanyaan:   jawaban_peserta.IdPertanyaan,
		JawabanPeserta: jawaban_peserta.JawabanPeserta,
		TotalSoal:      total_soal,
		JumlahBenar:    jumlah_benar,
		Skor:           jawaban_peserta.Skor,
		User: User{
            Nama: jawaban_peserta.User.Nama,
        },
        Quiz: Quiz{
            Judul: jawaban_peserta.Quiz.Judul,
        },
	}

	
	return res
}

/*

"message":"success",
"data": {
	"id": 1,
	"id_user": 1,
	nama: "nama",
	id_quiz: 1,
	title_quiz: "title",
	jumlah_benar: 9,
	total_soal: 10,
	skor: 90(misal admin memberi 90 maka artinya 90/9 yang mana pada tabel kolom skor akan terupdate dengan nilai 10 untuk yang benar dan 0 untuk yang salah)
}


*/