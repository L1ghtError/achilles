package storage

import (
	"ahls_srvi/internal/models"
	"sync"
	"time"
)

type Cel[T any] struct {
	expiresAt int64
	Data      T
}
type ACache struct {
	TTL       int64
	upstream  ICustodian
	soMtx     sync.RWMutex
	caMtx     sync.RWMutex
	crMtx     sync.RWMutex
	Sources   map[uint64]Cel[models.Source]
	Campaigns map[uint64]Cel[models.Campaign]
	Creatives map[uint64]Cel[models.Creative]
}

func NewACache(ttl int64, upstreamStorage ICustodian) ACache {
	return ACache{
		TTL:       ttl,
		upstream:  upstreamStorage,
		Sources:   make(map[uint64]Cel[models.Source]),
		Campaigns: make(map[uint64]Cel[models.Campaign]),
		Creatives: make(map[uint64]Cel[models.Creative]),
	}
}

func filterEntitiesBatch[T models.Source | models.Campaign | models.Creative](now int64, rw *sync.RWMutex, ids []uint64, ent map[uint64]Cel[T]) ([]T, []uint64) {
	out := make([]T, 0, len(ids))
	missing := make([]uint64, 0, len(ids))
	rw.RLock()
	defer rw.RUnlock()
	for _, id := range ids {
		if item, ok := ent[id]; ok && now < item.expiresAt {
			out = append(out, item.Data)
		} else {
			missing = append(missing, id)
		}
	}
	return out, missing
}

func (c *ACache) GetSource(id uint64) (models.Source, bool) {
	now := time.Now().Unix()
	if item, ok := c.Sources[id]; ok && now < item.expiresAt {
		c.soMtx.RLock()
		defer c.soMtx.RUnlock()
		return item.Data, ok
	} else {
		c.soMtx.Lock()
		defer c.soMtx.Unlock()
		v, ok := c.upstream.GetSource(id)
		c.Sources[id] = Cel[models.Source]{expiresAt: now + c.TTL, Data: v}
		return v, ok
	}
}

func (c *ACache) GetCampaignBatch(ids []uint64) []models.Campaign {
	now := time.Now().Unix()
	out, missing := filterEntitiesBatch(now, &c.caMtx, ids, c.Campaigns)
	if len(missing) != 0 {
		f := c.upstream.GetCampaignBatch(missing)
		out = append(out, f...)

		c.caMtx.Lock()
		defer c.caMtx.Unlock()
		for _, v := range f {
			c.Campaigns[v.ID] = Cel[models.Campaign]{expiresAt: now + c.TTL, Data: v}
		}
	}
	return out
}

func (c *ACache) GetCreativeBatch(ids []uint64) []models.Creative {
	now := time.Now().Unix()
	out, missing := filterEntitiesBatch(now, &c.crMtx, ids, c.Creatives)
	if len(missing) != 0 {
		f := c.upstream.GetCreativeBatch(missing)
		out = append(out, f...)

		c.crMtx.Lock()
		defer c.crMtx.Unlock()
		for _, v := range f {
			c.Creatives[v.ID] = Cel[models.Creative]{expiresAt: now + c.TTL, Data: v}
		}
	}
	return out
}
