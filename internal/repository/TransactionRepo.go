package repository

import (
	"errors"

	"azevedoruan.github/lbd-gorm/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

type MusicTransferRequest struct {
	MusicId        uint
	FromPlaylistId uint
	ToPlaylistId   uint
	UserId         uint
}

func (r *TransactionRepository) TransferMusicBetweenPlaylists(request MusicTransferRequest) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var fromPlaylist models.Playlist
		if err := tx.
			Where("playlist_id = ? AND usuario_id = ?", request.FromPlaylistId, request.UserId).
			First(&fromPlaylist).Error; err != nil {
			return err
		}

		var toPlaylist models.Playlist
		if err := tx.
			Where("playlist_id = ? AND usuario_id = ?", request.ToPlaylistId, request.UserId).
			First(&toPlaylist).Error; err != nil {
			return err
		}

		var sourceMusic models.MusicPlaylist
		if err := tx.
			Where(
				"musica_id = ? AND playlist_id = ? AND usuario_id = ?",
				request.MusicId,
				request.FromPlaylistId,
				request.UserId,
			).
			First(&sourceMusic).Error; err != nil {
			return err
		}

		var existingDestination models.MusicPlaylist
		err := tx.
			Where(
				"musica_id = ? AND playlist_id = ? AND usuario_id = ?",
				request.MusicId,
				request.ToPlaylistId,
				request.UserId,
			).
			First(&existingDestination).Error
		if err == nil {
			return errors.New("music already exists in destination playlist")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		var lastOrder uint
		if err := tx.Model(&models.MusicPlaylist{}).
			Select("COALESCE(MAX(ordem_na_playlist), 0)").
			Where("playlist_id = ? AND usuario_id = ?", request.ToPlaylistId, request.UserId).
			Scan(&lastOrder).Error; err != nil {
			return err
		}

		if err := tx.Delete(&sourceMusic).Error; err != nil {
			return err
		}

		destinationMusic := models.MusicPlaylist{
			MusicId:       request.MusicId,
			PlaylistId:    request.ToPlaylistId,
			UserId:        request.UserId,
			PlaylistOrder: lastOrder + 1,
		}
		if err := tx.Create(&destinationMusic).Error; err != nil {
			return err
		}

		return nil
	})
}
