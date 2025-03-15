package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID   uint   `json:"user_id" gorm:"index"` // Hapus `uniqueIndex`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Phone    string `json:"phone"`
	Image    string `json:"image"`
}
