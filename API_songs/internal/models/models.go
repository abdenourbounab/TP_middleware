package models

import (
	"github.com/gofrs/uuid"
)

type Songs struct {
	Song_id   *uuid.UUID `json:"song_id"`
	Album     string     `json:"album"`
	Artist    string     `json:"artist"`
	Song_name string     `json:"song_name"`
	Duration  float32    `json:"duration"`
	Type      string     `json:"type_song"`
	Playlist  string     `json:"playlist"`
}
