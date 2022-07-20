package loader

import "github.com/bwrega/btdogg-bff-go/internal/details/dto"
import _ "embed"
import "encoding/json"
import "fmt"

//go:embed torrents.json
var fileContent []byte

func Load() []dto.MongoTorrent {
	var torrents []dto.MongoTorrent
	if err := json.Unmarshal(fileContent, &torrents) ; err != nil {
		panic(fmt.Sprint("Failed to load mock torrents ", err))
	}
	return torrents
}
