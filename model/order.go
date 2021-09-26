package model

import "time"

type Order struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	Id_barang_order int       `json:"id_barang_order"`
	Barang_nama     string    `json:"barang_nama"`
	User_order      int       `json:"user_order"`
	Quantity_order  int       `json:"quantity_order"`
	Total_order     int       `json:"total_order"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
