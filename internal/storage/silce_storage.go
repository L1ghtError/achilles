package storage

import (
	"ahls_srvi/internal/models"
	"time"
)

type SliceStorage struct {
	Sources   []models.Source
	Campaigns []models.Campaign
	Creatives []models.Creative
}

var Mock1 SliceStorage = SliceStorage{
	Sources: []models.Source{
		{ID: 0, Name: "main", Status: models.Active, Campaigns: []uint64{0}},
		{ID: 1, Name: "crypto", Status: models.Active, Campaigns: []uint64{1}},
		{ID: 2, Name: "investment", Status: models.Active, Campaigns: []uint64{1}},
	},
	Campaigns: []models.Campaign{
		{ID: 0, Name: "self promotion", StartTime: time.Now(), EndTime: time.Now().Add(time.Hour * 10), Sources: []uint64{0}, Creatives: []uint64{0, 3, 4}},
		{ID: 1, Name: "fancy videos", StartTime: time.Now(), EndTime: time.Now().Add(time.Hour * 24), Sources: []uint64{1, 2}, Creatives: []uint64{1, 2, 5, 6, 7}},
	},
	Creatives: []models.Creative{
		{ID: 0, Price: 91_000, Duration: 12000, Playlist: ("Blank 0")},
		{ID: 1, Price: 490_000, Duration: 32000, Playlist: ("Blank 1")},
		{ID: 2, Price: 76_000, Duration: 37000, Playlist: ("Blank 2")},
		{ID: 3, Price: 210_000, Duration: 44000, Playlist: ("Blank 3")},
		{ID: 4, Price: 115_000, Duration: 23000, Playlist: ("Blank 4")},
		{ID: 5, Price: 131_000, Duration: 25000, Playlist: ("Blank 5")},
		{ID: 6, Price: 41_000, Duration: 17000, Playlist: ("Blank 6")},
		{ID: 7, Price: 74_000, Duration: 22000, Playlist: ("Blank 7")},
	},
}

func (ss *SliceStorage) GetSource(id uint64) (models.Source, bool) {
	for i := 0; i < len(ss.Sources); i++ {
		if id == ss.Sources[i].ID {
			return ss.Sources[i], true
		}
	}
	return models.Source{}, false
}

func (ss *SliceStorage) GetCampaignBatch(ids []uint64) []models.Campaign {
	matched := make(map[uint64]int, len(ids))
	c := make([]models.Campaign, 0, len(ids))
	for i := 0; i < len(ss.Campaigns); i++ {
		matched[ss.Campaigns[i].ID] = i
	}
	for _, id := range ids {
		if v, ok := matched[id]; ok {
			c = append(c, ss.Campaigns[v])
		}
	}
	return c
}

func (ss *SliceStorage) GetCreativeBatch(ids []uint64) []models.Creative {
	matched := make(map[uint64]int, len(ids))
	c := make([]models.Creative, 0, len(ids))
	for i := 0; i < len(ss.Creatives); i++ {
		matched[ss.Creatives[i].ID] = i
	}
	for _, id := range ids {
		if v, ok := matched[id]; ok {
			c = append(c, ss.Creatives[v])
		}
	}
	return c
}
