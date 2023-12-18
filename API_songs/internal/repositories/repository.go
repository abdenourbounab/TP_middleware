package songs

import (
	"github.com/gofrs/uuid"
	"songs/internal/helpers"
	"songs/internal/models"
)

func GetAllSongs() ([]models.Songs, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Songs{}
	for rows.Next() {
		var data models.Songs
		err = rows.Scan(&data.Song_id, &data.Album, &data.Artist, &data.Artist, &data.Song_name, &data.Duration, &data.Type, &data.Playlist)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Songs, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Songs
	err = row.Scan(&song.Song_id, &song.Artist, &song.Song_name)
	if err != nil {
		return nil, err
	}
	return &song, err
}
