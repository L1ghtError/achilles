package models

import "time"

type Status bool

const (
	Active   Status = true
	Inactive Status = false
)

type Source struct {
	ID        uint64   `json:"ID"`
	Name      string   `json:"Name"`
	Status    Status   `json:"Status"`
	Campaigns []uint64 `json:"Campaigns"`
}

type Campaign struct {
	ID        uint64    `json:"ID"`
	Name      string    `json:"Name"`
	StartTime time.Time `json:"StartTime"`
	EndTime   time.Time `json:"EndTime"`
	Sources   []uint64  `json:"Sources"`
	Creatives []uint64  `json:"Creatives"`
}

type Creative struct {
	ID       uint64 `json:"ID"`
	Price    uint64 `json:"Price"`    // Price in cents
	Duration uint64 `json:"Duration"` // Time in miliseconds
	Playlist string `json:"Playlist"`
}
