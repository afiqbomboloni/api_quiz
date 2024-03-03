package response

import "github.com/afiqbomboloni/api_quiz/entity"

type QuizResponse struct {
	ID           uint   `json:"id"`
	Judul        string `json:"judul"`
	Deskripsi    string `json:"deskripsi"`
	WaktuMulai   string `json:"waktu_mulai"`
	WaktuSelesai string `json:"waktu_selesai"`
	TotalSoal    int64  `json:"total_soal"`
	Pertanyaan    []struct {
		ID           uint   `json:"id"`
		Pertanyaan   string `json:"pertanyaan"`
		OpsiJawaban  string `json:"opsi_jawaban"`
		JawabanBenar int `json:"jawaban_benar"`
	} `json:"questions"`
}

func NewQuizResponse(quiz entity.Quiz, total_soal int64) QuizResponse {
	res := QuizResponse{
		ID:           quiz.ID,
		Judul:        quiz.Judul,
		Deskripsi:    quiz.Deskripsi,
		WaktuMulai:   quiz.WaktuMulai.Format("2006-01-02T15:04:05Z"),
		WaktuSelesai: quiz.WaktuSelesai.Format("2006-01-02T15:04:05Z"),
		TotalSoal:    total_soal,
		Pertanyaan:    []struct {
			ID           uint   `json:"id"`
			Pertanyaan   string `json:"pertanyaan"`
			OpsiJawaban  string `json:"opsi_jawaban"`
			JawabanBenar int `json:"jawaban_benar"`
		}{},
	}

	for _, q := range quiz.Pertanyaan {
		res.Pertanyaan = append(res.Pertanyaan, struct {
			ID           uint   `json:"id"`
			Pertanyaan   string `json:"pertanyaan"`
			OpsiJawaban  string `json:"opsi_jawaban"`
			JawabanBenar int `json:"jawaban_benar"`
		}{
			ID:           q.ID,
			Pertanyaan:   q.Pertanyaan,
			OpsiJawaban:  q.OpsiJawaban,
			JawabanBenar: q.JawabanBenar,
		})
	}

	return res
}
