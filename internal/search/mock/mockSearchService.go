package mock

import (
	"github.com/bwrega/btdogg-bff-go/internal/details"
	"github.com/bwrega/btdogg-bff-go/internal/details/mock"
	"github.com/bwrega/btdogg-bff-go/internal/details/mock/loader"
	"github.com/bwrega/btdogg-bff-go/internal/search"
	"strings"
)

type MockSearchService struct {
	torrents []details.TorrentDetails
}

func NewMockSearchService() search.TorrentSearchService {
	mongoTorrents := loader.Load()
	return MockSearchService{torrents: mock.FlatTorrents(mongoTorrents)}
}


func (mss MockSearchService) Search(query string) search.ResultList {
	if query == "" {
		return search.ResultList{
			Items: []search.ResultItem {},
			TotalResults: 0,
		}
	}
	var matchedTorrents []details.TorrentDetails
	for _, torrent := range mss.torrents {
		if matches(torrent, query) {
			matchedTorrents = append(matchedTorrents, torrent)
		}
	}
	resultItems := toResultItems(matchedTorrents)
	return search.ResultList{
		Items: resultItems,
		TotalResults: len(resultItems),
	}
}

func matches(torrent details.TorrentDetails, query string) bool {
	query = strings.ToLower(query)
	if strings.Contains(strings.ToLower(torrent.Title), query) {
		return true
	}
	for _, file := range torrent.Files {
		if strings.Contains(strings.ToLower(file.Name), query) {
			return true
		}
	}
	return false
}

func toResultItems(torrents []details.TorrentDetails) []search.ResultItem {
	results := []search.ResultItem{}
	for i, torrent := range torrents {
		results = append(results, search.ResultItem{
			Id: torrent.Id,
			Title: torrent.Title,
			TotalSize: torrent.TotalSize,
			Liveness: (i/3)*(i/3),
			Magnet: torrent.Magnet,
			CreatedDate: torrent.CreatedDate,
		})
	}
	return results
}
