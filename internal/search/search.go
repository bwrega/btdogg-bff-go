package search

import (
	"github.com/bwrega/btdogg-bff-go/internal/details"
	"time"
)

type ResultList struct {
	Items []ResultItem `json:"items"`
	TotalResults int   `json:"totalResults"`
}

type ResultItem struct {
	Id details.TorrentHash `json:"id"`
	Title string `json:"title"`
	TotalSize uint64 `json:"totalSize"`
	Liveness int `json:"liveness"`
	Magnet string `json:"magnet"`
	CreatedDate time.Time `json:"createdDate"`
}

type TorrentSearchService interface {
	Search(query string) ResultList
}
