package models

type Music struct {
	ID          uint   `gorm:"column:id;primaryKey"`
	Title       string `gorm:"column:titulo"`
	DurationSec int32  `gorm:"column:duracao_segundos"`
	ArtistId    uint   `gorm:"column:artista_id"`
	Artist      Artist `gorm:"foreignKey:ArtistId"`
}

func (Music) TableName() string {
	return "musica"
}
