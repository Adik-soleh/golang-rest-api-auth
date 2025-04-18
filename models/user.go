package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string   `json:"name"`
	Email    string   `json:"email" gorm:"unique"`
	Password string   `json:"-"`
	Profile  *Profile `json:"profile" gorm:"constraint:OnDelete:CASCADE;"`
}
