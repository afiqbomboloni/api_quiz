package entity

import "golang.org/x/crypto/bcrypt"

// import (
// 	"github.com/google/uuid"
// )

type User struct {
    // ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique_index"`
    ID             uint             `gorm:"primaryKey"`
    Nama           string           `gorm:"size:255;not null;unique" json:"nama"`
    Password       string           `gorm:"size:255;not null;" json:"password"`
    Email          string           `gorm:"size:100;not null;unique" json:"email"`
    Role           string           `gorm:"default:'user';size:255;not null;" json:"role"`
    JawabanPeserta []JawabanPeserta `gorm:"foreignKey:IdUser"`
}

func NewUser(nama, email, password, role string) *User {
    newPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return &User{
        Nama:     nama,
        Email:    email,
        Password: string(newPassword),
        Role:     role,
    }
}