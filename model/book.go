package model

import "time"

type Book struct {
	ID        uint      `gorm:"not null;primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"not null;unique;varchar(255)" json:"name_book"`
	Author    string    `gorm:"not null;varchar(255)" json:"author"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
