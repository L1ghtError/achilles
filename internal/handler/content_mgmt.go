package handler

import (
	"ahls_srvi/internal/models"
	"ahls_srvi/internal/storage"
	"net/http"
	"strconv"
	"strings"
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
		http.Error(w, "Warning, Err in Query", http.StatusBadRequest)
		return
	}

	maxDuration := r.URL.Query().Get("maxDuration") // Get query param
	maxDur, err := strconv.Atoi(maxDuration)
	if err != nil {
		http.Error(w, "Warning, Err in Query", http.StatusBadRequest)
		return
	}

	ss := ah.Cache
	source, ok := ss.GetSource(uint64(srcID))
	if !ok {
		w.Write([]byte("Warning, Err: " + sourceID + " " + maxDuration))
		return
	}
	if source.Status == models.Inactive {
		http.Error(w, "Status inactive: "+sourceID+" "+maxDuration, http.StatusGone)
		return
	}
	campaigns := ss.GetCampaignBatch(source.Campaigns)
	if len(campaigns) == 0 {
		http.Error(w, "There is 0 campaigns: "+sourceID+" "+maxDuration, http.StatusNotFound)
		return
	}

	cids := creativityFromCampain(campaigns)
	if len(cids) == 0 {
		http.Error(w, "There is no suitable creativities: "+sourceID+" "+maxDuration, http.StatusNotFound)
		return
	}
	creatives := ss.GetCreativeBatch(cids)
	if len(creatives) == 0 {
		http.Error(w, "There is no suitable creativities: "+sourceID+" "+maxDuration, http.StatusNotFound)
		return
	}
	ci := appropriateCreativityIndex(creatives, maxDur)
	if ci == -1 {
		http.Error(w, "There is no suitable creativities: "+sourceID+" "+maxDuration, http.StatusNotFound)
		return
	}
	id := strconv.Itoa(int(creatives[ci].ID))
	w.Header().Set("X-Resource-ID", id)
	w.Write([]byte(strings.ReplaceAll(creatives[ci].Playlist, "\\n", "\n")))
}
