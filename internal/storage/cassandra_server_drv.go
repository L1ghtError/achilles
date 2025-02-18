package storage

import (
	"ahls_srvi/internal/models"
	"ahls_srvi/internal/strutil"
	"context"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

type CassandraServerOps struct {
	Hostnames []string
	Keyspace  string
}

type CassandraServerDrv struct {
	ops     CassandraServerOps
	cluster *gocql.ClusterConfig
	session *gocql.Session
	timeout time.Duration
}

func NewCassandraServerDrv(ops CassandraServerOps) (*CassandraServerDrv, func()) {
	cluster := gocql.NewCluster(ops.Hostnames...)
	cluster.Keyspace = ops.Keyspace
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, nil
	}
	tm := time.Second * 10
	return &CassandraServerDrv{ops: ops, cluster: cluster, session: session, timeout: tm}, session.Close
}

func (c *CassandraServerDrv) GetSource(id uint64) (models.Source, bool) {
	ctx, _ := context.WithTimeout(context.Background(), c.timeout)

	iter := c.session.Query(`SELECT id,name,status,campaign FROM source_by_id  WHERE id = ?`,
		strconv.Itoa(int(id))).WithContext(ctx).Iter()
	defer iter.Close()
	src := models.Source{Campaigns: make([]uint64, 0, iter.NumRows())}
	var campaign uint64

	if iter.Scan(&src.ID, &src.Name, &src.Status, &campaign) {
		src.Campaigns = append(src.Campaigns, campaign)
		scanner := iter.Scanner()
		for scanner.Next() {
			err := scanner.Scan(nil, nil, nil, &campaign)
			if err != nil {
				return models.Source{}, false
			}
			src.Campaigns = append(src.Campaigns, campaign)
		}
	}

	return src, true
}

func (c *CassandraServerDrv) GetCampaignBatch(ids []uint64) []models.Campaign {
	if len(ids) < 1 {
		return []models.Campaign{}
	}
	ctx, _ := context.WithTimeout(context.Background(), c.timeout)
	ca := make([]models.Campaign, 0, len(ids))
	query := `SELECT id, name, start_time, end_time, creative 
	FROM campaign_by_id WHERE id IN (` + strutil.JoinInts(ids, ",") + `)`

	scanner := c.session.Query(query).WithContext(ctx).Iter().Scanner()

	for scanner.Next() {
		var c models.Campaign
		err := scanner.Scan(&c.ID, &c.Name, &c.StartTime, &c.EndTime, &c.Creatives)
		if err != nil {
			return []models.Campaign{}
		}
		ca = append(ca, c)
	}

	return ca
}

func (c *CassandraServerDrv) GetCreativeBatch(ids []uint64) []models.Creative {
	if len(ids) < 1 {
		return []models.Creative{}
	}
	ctx, _ := context.WithTimeout(context.Background(), c.timeout)
	cr := make([]models.Creative, 0, len(ids))
	query := `SELECT id, price, duration, playlist 
	FROM creative_by_id WHERE id IN (` + strutil.JoinInts(ids, ",") + `)`

	scanner := c.session.Query(query).WithContext(ctx).Iter().Scanner()

	for scanner.Next() {
		var c models.Creative
		err := scanner.Scan(&c.ID, &c.Price, &c.Duration, &c.Playlist)
		if err != nil {
			return []models.Creative{}
		}
		cr = append(cr, c)
	}

	return cr
}
