package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type Status struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Link   string `json:"link"`
	UserID int    `json:"user_id" gorm:"foreignKey:User;references:ID"`
}
