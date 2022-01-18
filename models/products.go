package models

import "time"

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" `
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// TableName gives table name of model
func (p Product) TableName() string {
	return "products"
}
