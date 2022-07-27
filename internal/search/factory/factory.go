package factory

import (
	"github.com/bwrega/btdogg-bff-go/internal/search"
	"github.com/bwrega/btdogg-bff-go/internal/search/mock"
)

func NewTorrentSearchService(kind TorrentSearchServiceType) search.TorrentSearchService {
	switch kind {
	case Mock:
		return mock.NewMockSearchService()
	case Real:
		return nil
	default:
		panic("Unknown TorrentSearchServiceType: " + kind)
	}
}

type TorrentSearchServiceType string

const (
	Real TorrentSearchServiceType = "Real"
	Mock TorrentSearchServiceType = "Mock"
)
