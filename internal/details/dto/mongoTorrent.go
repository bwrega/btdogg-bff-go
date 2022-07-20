package dto

import (
	"github.com/bwrega/btdogg-bff-go/internal/details"
	"time"
	"cloud.google.com/go/civil"
)

type MongoTorrent struct {
	Id details.TorrentHash
	Title string
	TotalSize uint64
	Data []TorrentFileOrDirectory
	Creation time.Time
	Modification time.Time
	Liveness Liveness
}

type TorrentFileOrDirectory struct {
	Name string
	Size uint64
	Contents []TorrentFileOrDirectory
}

func (tf TorrentFileOrDirectory) isFile() bool {
	return tf.Contents == nil
}

func (tf TorrentFileOrDirectory) isDir() bool {
	return !tf.isFile()
}

type Liveness struct {
	Requests map[civil.Date]int
	Announces map[civil.Date]int
}
