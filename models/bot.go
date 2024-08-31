package models

type Bot struct {
	Id int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Ip string `json:"ip"`
}
