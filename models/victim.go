package models

type Victim struct {
	Id int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Ip string `json:"ip"`
}
