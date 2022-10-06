package models

type Car struct {
	Merk     string `gorm:"varchar(100)"`
	Harga    int    `gorm:"integer(11)"`
	Typecars string `gorm:"varchar(100)"`
	Id       uint   `gorm:"primaryKey"`
}
