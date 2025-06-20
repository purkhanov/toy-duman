package model

type User struct {
	ID       int      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string   `json:"username"`
	Name     string   `json:"name"`
	Status   []Status `json:"status" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type Status struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Link   string `json:"link"`
	UserID int    `json:"user_id"`
}
