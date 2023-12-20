package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"songs/internal/controllers/songs"
	"songs/internal/helpers"
	_ "songs/internal/models"
)

func main() {
	r := chi.NewRouter()

	r.Post("/songs", songs.InsertSong)
	r.Get("/songs", songs.GetSongs)
	r.Get("/songs/{id}", songs.GetSong)
	r.Put("/songs/{id}", songs.UpdateSong)

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS songs (
			song_id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			album VARC HAR(255) NOT NULL,
    		artist VARCHAR(255) NOT NULL,
    		song_name VARCHAR(255) NOT NULL,
    		duration FLOAT NOT NULL,
    		type_song VARCHAR(255) NOT NULL,
    		playlist VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
