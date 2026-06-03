package models

type Artist struct {
	ID          uint   `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:nome"`
	Nacionality string `gorm:"column:nacionalidade"`
}

func (Artist) TableName() string {
	return "artista"
}
