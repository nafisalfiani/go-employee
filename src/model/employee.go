package model

type Employee struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name     string `json:"name" gorm:"column:name"`
	TeamName string `json:"team_name" gorm:"column:team_name"`
}
