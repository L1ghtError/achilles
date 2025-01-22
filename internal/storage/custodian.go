package storage

import "ahls_srvi/internal/models"

type ICustodian interface {
	GetSource(id uint64) (models.Source, bool)
	GetCampaignBatch(ids []uint64) []models.Campaign
	GetCreativeBatch(ids []uint64) []models.Creative
}
