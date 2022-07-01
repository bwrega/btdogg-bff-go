package mock

import "github.com/bwrega/btdogg-bff-go/internal/details"

type MockTorrentDetailsService struct {}

func NewMockTorrentDetailsService() details.TorrentDetailsService {
	return MockTorrentDetailsService{}
}

func (s MockTorrentDetailsService) GetDetails(hash details.TorrentHash) *details.TorrentDetails  {
	return nil
}
