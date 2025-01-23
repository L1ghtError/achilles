package handler

import (
	"ahls_srvi/internal/m3u8_parser"
	"ahls_srvi/internal/queries"
	"io"
	"net/http"
	"strconv"
)

type AdHandler struct {
	C queries.IClient
}

func (ah *AdHandler) StichM3U8(w http.ResponseWriter, r *http.Request) {
	sourceID := r.URL.Query().Get("sourceID")
	srcID, err := strconv.Atoi(sourceID)
	if err != nil {
		http.Error(w, "Error, bad srcID", http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}
	adSegments := string(body)

	duration := int(m3u8_parser.CountTotalINF(adSegments)) + 1
	playlist, id := ah.C.GetRecommendedPlaylist(srcID, duration)
	newPlaylist, isModified := m3u8_parser.ModifyPlaylist(playlist, adSegments)
	if !isModified {
		http.Error(w, "Failed to modify playlist", http.StatusInternalServerError)
		return
	}
	success := ah.C.UpdatePlaylist(newPlaylist, id)
	if !success {
		http.Error(w, "Failed to update playlist", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(newPlaylist))
}
