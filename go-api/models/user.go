package models

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	AccountID uint   `json:"account_id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
}