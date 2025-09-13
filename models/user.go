package models

type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Age   int    `gorm:"not null" json:"age"`
	Name  string `gorm:"size:100;not null" json:"name"`
	Email string `gorm:"size:100;unique;not null" json:"email"`
}
