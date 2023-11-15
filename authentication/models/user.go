package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(300)" json:"username"`
	Password string `gorm:"type:varchar(300)" json:"password"`
}
