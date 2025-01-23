package queries

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type LClient struct {
	Hostname string
	DBaddr   string
}

func (c *LClient) GetRecommendedPlaylist(srcId int, maxDuration int) (string, int) {
	url := fmt.Sprintf("%s/auction?sourceID=%d&maxDuration=%d", c.Hostname, srcId, maxDuration)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0
	}
	resourceID := resp.Header.Get("X-Resource-ID")
	id, _ := strconv.Atoi(resourceID)
	return string(body), id

}

func (c *LClient) UpdatePlaylist(pNew string, creativeId int) bool {
	url := fmt.Sprintf("%s/Creatives/%d", c.DBaddr, creativeId)

	body := struct{ Playlist string }{Playlist: pNew}
	jbody, err := json.Marshal(body)
	if err != nil {
		return false
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jbody))
	if err != nil {
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	_ = resp

	return err == nil
}
