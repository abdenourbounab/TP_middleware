package models

import (
	"github.com/gofrs/uuid"
)

type Songs struct {
	Id       *uuid.UUID `json:"id"`
	Album    string     `json:"album"`
	Artist   string     `json:"artist"`
	SongName string     `json:"song_name"`
	Duration float32    `json:"duration"`
	Type     string     `json:"type_song"`
	Playlist string     `json:"playlist"`
}
