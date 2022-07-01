package details

import (
	"time"
)

type TorrentDetails struct {
	Id TorrentHash `json:"id"`
	Title string `json:"title"`
	TotalSize uint64 `json:"totalSize"`
	Files []TorrentFile `json:"files"`
	CreatedDate time.Time `json:"createdDate"`
	Liveness int `json:"liveness"`
	Magnet string `json:"magnet"`
}

type TorrentFile struct {
	Name string `json:"name"`
	Size uint64 `json:"size"`
}

type TorrentHash string

type TorrentDetailsService interface {
	GetDetails(hash TorrentHash) *TorrentDetails
}
