package models

import "gorm.io/gorm"

type Users struct{
	gorm.Model
	Name string `json:"name"`
	Password string `json:"password"`
}