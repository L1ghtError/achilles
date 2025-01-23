package handler

import (
	"ahls_srvi/internal/models"
	"ahls_srvi/internal/storage"
	"net/http"
	"strconv"
	"time"
)

type AppHandler struct {
	Cache storage.ICustodian
}

func creativityFromCampain(campaigns []models.Campaign) []uint64 {
	cid := 0
	now := time.Now()
	for _, v := range campaigns {
		cid += len(v.Creatives)
	}
	cids := make([]uint64, 0, cid)

	for _, v := range campaigns {
		if v.EndTime.After(now) {
			cids = append(cids, v.Creatives...)
		}
	}
	return cids
}

func appropriateCreativityIndex(creatives []models.Creative, maxDur int) int {
	var maxPrice uint64 = 0
	b := -1
	for i, v := range creatives {
		if v.Duration <= uint64(maxDur) && maxPrice < v.Price {
			maxPrice = v.Price
			b = i
		}
	}
	return b
}

func (ah *AppHandler) GetAppropriateContent(w http.ResponseWriter, r *http.Request) {
	sourceID := r.URL.Query().Get("sourceID") // Get query param
	srcID, err := strconv.Atoi(sourceID)
	if err != nil {
		// !! TODO: handle err
		w.Write([]byte("Warning, Err in Query"))
		return
	}

	maxDuration := r.URL.Query().Get("maxDuration") // Get query param
	maxDur, err := strconv.Atoi(maxDuration)
	if err != nil {
		// !! TODO: handle err
		w.Write([]byte("Warning, Err in Query"))
		return
	}

	ss := ah.Cache
	source, ok := ss.GetSource(uint64(srcID))
	if !ok {
		w.Write([]byte("Warning, Err: " + sourceID + " " + maxDuration))
		return
	}
	if source.Status == models.Inactive {
		w.Write([]byte("Status inactive: " + sourceID + " " + maxDuration))
		return
	}
	campaigns := ss.GetCampaignBatch(source.Campaigns)
	if len(campaigns) == 0 {
		w.Write([]byte("There is 0 campaigns: " + sourceID + " " + maxDuration))
		return
	}

	cids := creativityFromCampain(campaigns)
	if len(cids) == 0 {
		w.Write([]byte("There is no suitable creativities: " + sourceID + " " + maxDuration))
		return
	}
	creatives := ss.GetCreativeBatch(cids)
	if len(creatives) == 0 {
		w.Write([]byte("There is no suitable creativities: " + sourceID + " " + maxDuration))
		return
	}
	ci := appropriateCreativityIndex(creatives, maxDur)
	if ci == -1 {
		w.Write([]byte("There is no suitable creativities: " + sourceID + " " + maxDuration))
		return
	}
	id := strconv.Itoa(int(creatives[ci].ID))
	w.Header().Set("X-Resource-ID", id)
	w.Write([]byte(creatives[ci].Playlist))
}
