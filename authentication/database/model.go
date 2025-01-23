package database


type User struct {
	UserID     uint   `gorm:"primaryKey" json:"user_id"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
}

