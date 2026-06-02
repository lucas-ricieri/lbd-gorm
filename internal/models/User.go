package models

type User struct {
	ID    uint   `gorm:"column:id;primaryKey"`
	Name  string `gorm:"column:username"`
	Email string `gorm:"column:email"`
}

func (User) TableName() string {
	return "usuario"
}
