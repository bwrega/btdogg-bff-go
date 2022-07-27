package mock

import (
	"github.com/bwrega/btdogg-bff-go/internal/details"
	"github.com/bwrega/btdogg-bff-go/internal/details/dto"
	"github.com/bwrega/btdogg-bff-go/internal/details/mock/loader"
)

type MockTorrentDetailsService struct {
	torrents []details.TorrentDetails
}

func NewMockTorrentDetailsService() details.TorrentDetailsService {
	mongoTorrents := loader.Load()
	return MockTorrentDetailsService{torrents: FlatTorrents(mongoTorrents)}
}

func FlatTorrents(mongoTorrents []dto.MongoTorrent) []details.TorrentDetails {
	var torrentsDetails []details.TorrentDetails
	for _, mt := range mongoTorrents {
		torrentsDetails = append(torrentsDetails, mt.ToDetails())
	}
	return torrentsDetails
}

func (s MockTorrentDetailsService) GetDetails(hash details.TorrentHash) *details.TorrentDetails  {
	for _, t := range s.torrents {
		if t.Id == hash {
			return &t
		}
	}
	return nil
}
