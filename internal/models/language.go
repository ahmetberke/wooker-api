package models

type Language struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
	NativeName string `json:"native_name"`
}