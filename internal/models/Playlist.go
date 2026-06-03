package models

import "time"

type Playlist struct {
	PlaylistId    uint      `gorm:"column:playlist_id;primaryKey"`
	UserId        uint      `gorm:"column:usuario_id;primaryKey"`
	Name          uint      `gorm:"column:nome"`
	CreactionDate time.Time `gorm:"column:data_criacao"`
}
