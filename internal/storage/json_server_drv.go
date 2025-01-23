package storage

import (
	"ahls_srvi/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type JsonServerDrv struct {
	Hostname string
}

func (c *JsonServerDrv) GetSource(id uint64) (models.Source, bool) {
	url := fmt.Sprintf("%s/Sources/%d", c.Hostname, id)
	resp, err := http.Get(url)
	if err != nil {
		return models.Source{}, false
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Source{}, false
	}

	var src models.Source
	err = json.Unmarshal([]byte(body), &src)
	if err != nil {
		log.Print(err)
		return models.Source{}, false
	}
	return src, true
}

func (c *JsonServerDrv) GetCampaignBatch(ids []uint64) []models.Campaign {
	query := ""
	for _, v := range ids {
		query += "id=" + strconv.Itoa(int(v)) + "&"
	}
	url := fmt.Sprintf("%s/Campaigns?%s", c.Hostname, query)
	resp, err := http.Get(url)
	if err != nil {
		return []models.Campaign{}
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.Campaign{}
	}

	var ca []models.Campaign
	err = json.Unmarshal([]byte(body), &ca)
	if err != nil {
		log.Print(err)
		return []models.Campaign{}
	}
	return ca
}

func (c *JsonServerDrv) GetCreativeBatch(ids []uint64) []models.Creative {
	query := ""
	for _, v := range ids {
		query += "id=" + strconv.Itoa(int(v)) + "&"
	}
	url := fmt.Sprintf("%s/Creatives?%s", c.Hostname, query)
	resp, err := http.Get(url)
	if err != nil {
		return []models.Creative{}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.Creative{}
	}

	var cr []models.Creative
	err = json.Unmarshal([]byte(body), &cr)
	if err != nil {
		log.Print("Unmarshal err: ", err)
		return []models.Creative{}
	}
	return cr
}
