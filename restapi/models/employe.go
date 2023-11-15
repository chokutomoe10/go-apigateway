package models

type Employe struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(200)" json:"name"`
	Role string `gorm:"type:varchar(200)" json:"role"`
}
