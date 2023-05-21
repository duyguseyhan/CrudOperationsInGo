package models

import "time"

type Customer struct {
	ID        uint      `gorm:"primarykey"`
	FirstName string    `gorm:"type:varchar(100);not null"`
	LastName  string    `gorm:"type:varchar(100);not null"`
	BirthDate time.Time `gorm:"not null"`
	Gender    string    `gorm:"type:varchar(6);not null"`
	Email     string    `gorm:"type:varchar(100);not null;unique"`
	Address   string    `gorm:"type:varchar(200)"`
}
