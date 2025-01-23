package queries

type IClient interface {
	GetRecommendedPlaylist(srcId int, maxDuration int) (string, int)
	UpdatePlaylist(pNew string, creativeId int) bool
}
