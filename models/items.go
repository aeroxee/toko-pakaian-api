package models

import "time"

// category
type ItemCategory struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:50;unique" json:"title"`
	Description *string   `gorm:"size:255" json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	Items       []Item    `gorm:"foreignKey:CategoryID" json:"items"`
}

// Item
type Item struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	CategoryID  int       `json:"category_id"`
	Name        string    `gorm:"size:50" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	Quality     float32   `gorm:"default:4.5" json:"quality"`
	Price       uint      `gorm:"default:0" json:"price"`
	Stock       uint      `gorm:"default:0" json:"stock"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
