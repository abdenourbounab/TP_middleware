package songs

import (
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"songs/internal/helpers"
	"songs/internal/models"
)

func InsertSong(song models.Songs) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO songs ( id ,album, artist, song_name, duration, type_song, playlist) VALUES (?, ?, ?, ?, ?, ?, ?)",
		song.Id, song.Album, song.Artist, song.SongName, song.Duration, song.Type, song.Playlist)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion du user dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

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
		err = rows.Scan(&data.Id, &data.Album, &data.Artist, &data.Artist, &data.SongName, &data.Duration, &data.Type, &data.Playlist)
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
	err = row.Scan(&song.Id, &song.Artist, &song.SongName)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func UpdateSong(song *models.Songs) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)
	_, err = db.Exec("UPDATE songs SET album=?, artist=?, song_name=?, duration=?, type_song=?, playlist=? WHERE id=?",
		song.Id, song.Album, song.Artist, song.SongName, song.Duration, song.Type, song.Playlist)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du user dans la base de données : %s", err.Error())
		return err
	}

	return nil
}
