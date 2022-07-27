package main

import (
	"github.com/bwrega/btdogg-bff-go/internal/details/factory"
	searchFactory "github.com/bwrega/btdogg-bff-go/internal/search/factory"
	. "github.com/bwrega/btdogg-bff-go/internal/server"
)

func main() {
	HttpServer{
		PortNumber: 2599,
		DetailsService: factory.NewTorrentDetailsService(factory.Mock),
		SearchService: searchFactory.NewTorrentSearchService(searchFactory.Mock),
	}.Start()
}
