package models

import "time"

type Playlist struct {
	PlaylistId    uint      `gorm:"column:playlist_id;primaryKey;autoIncrement"`
	UserId        uint      `gorm:"column:usuario_id;primaryKey"`
	Name          string    `gorm:"column:nome"`
	CreactionDate time.Time `gorm:"column:data_criacao"`

	// Preloads
	User           User            `gorm:"foreignKey:UserId"`
	MusicPlaylists []MusicPlaylist `gorm:"foreignKey:PlaylistId,UserId"`
}

func (Playlist) TableName() string {
	return "playlist"
}
