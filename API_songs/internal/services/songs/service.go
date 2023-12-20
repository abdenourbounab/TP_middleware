package songs

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"songs/internal/models"
	repository "songs/internal/repositories/songs"
)

func InsertSong(song models.Songs) error {
	err := repository.InsertSong(song)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du user : %s", err.Error())
		return err
	}
	return nil
}

func GetAllSongs() ([]models.Songs, error) {
	var err error
	// calling repository
	songs, err := repository.GetAllSongs()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}

func GetSongById(id uuid.UUID) (*models.Songs, error) {
	song, err := repository.GetSongById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

func UpdateSong(songID uuid.UUID, updatedSong models.Songs) error {
	song, err := repository.GetSongById(songID)
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération du song : %s", err.Error())
		return err
	}

	// Mettre à jour les champs nécessaires du commentaire récupéré avec les données du commentaire mis à jour
	song.Album = updatedSong.Album
	song.Artist = updatedSong.Artist
	song.SongName = updatedSong.SongName
	song.Duration = updatedSong.Duration
	song.Type = updatedSong.Type
	song.Playlist = updatedSong.Playlist

	err = repository.UpdateSong(song)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du song en base de données : %s", err.Error())
		return err
	}

	return nil
}

func DeleteSong(songID uuid.UUID) error {
	err := repository.DeleteSong(songID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du commentaire : %s", err.Error())
		return err
	}

	return nil
}
