package request

type AuthRequest struct {
	Nama string `json:"nama" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
	Role string `json:"role"`
}