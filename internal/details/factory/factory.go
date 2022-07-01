package factory

import (
	"github.com/bwrega/btdogg-bff-go/internal/details"
	"github.com/bwrega/btdogg-bff-go/internal/details/mock"
)

func NewTorrentDetailsService(kind TorrentDetailsServiceType) details.TorrentDetailsService {
	switch kind {
	case Mock:
		return mock.NewMockTorrentDetailsService()
	case Real:
		return nil
	default:
		panic("Unknown TorrentDetailsServiceType: " + kind)
	}
}

type TorrentDetailsServiceType string

const (
	Real TorrentDetailsServiceType = "Real"
	Mock TorrentDetailsServiceType = "Mock"
)
