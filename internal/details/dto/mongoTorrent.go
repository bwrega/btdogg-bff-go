package dto

import (
	"github.com/bwrega/btdogg-bff-go/internal/details"
	"time"
	"cloud.google.com/go/civil"
)

type MongoTorrent struct {
	Id details.TorrentHash `json:"_id"`
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

func (tf TorrentFileOrDirectory) IsFile() bool {
	return tf.Contents == nil
}

func (tf TorrentFileOrDirectory) IsDir() bool {
	return !tf.IsFile()
}

type Liveness struct {
	Requests map[civil.Date]int
	Announces map[civil.Date]int
}

func (t MongoTorrent) ToDetails() details.TorrentDetails {
	var files []details.TorrentFile
	for _, mf := range t.Data {
		files = append(files, mf.toDetails()...)
	}
	return details.TorrentDetails{
		Id: t.Id,
		Title: t.Title,
		TotalSize: t.TotalSize,
		Files: files,
		CreatedDate: t.Creation,
		Liveness: 42,
		Magnet: "",
	}
}

func (tf TorrentFileOrDirectory) toDetails() []details.TorrentFile {
	if tf.IsFile() {
		return []details.TorrentFile{{Name: tf.Name, Size: tf.Size}}
	} else {
		var ret []details.TorrentFile
		var innerFiles []details.TorrentFile
		for _, file := range tf.Contents {
			innerFiles = append(innerFiles, file.toDetails()...)
		}
		for _, t := range innerFiles {
			ret = append(ret, details.TorrentFile{Name: tf.Name + "/" + t.Name, Size: t.Size})
		}
		return ret
	}
}
