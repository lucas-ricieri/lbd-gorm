package models

type MusicPlaylist struct {
	MusicId       uint  `gorm:"column:musica_id;primaryKey"`
	PlaylistId    uint  `gorm:"column:playlist_id;primaryKey"`
	UserId        uint  `gorm:"column:usuario_id;primaryKey"`
	PlaylistOrder uint  `gorm:"column:ordem_na_playlist"`
	Music         Music `gorm:"foreignKey:MusicId"` // Preload
}

func (MusicPlaylist) TableName() string {
	return "musica_playlist"
}
