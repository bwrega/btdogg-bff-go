package main

import (
	"github.com/bwrega/btdogg-bff-go/internal/details/factory"
	. "github.com/bwrega/btdogg-bff-go/internal/server"
)

func main() {
	HttpServer{PortNumber: 2599, DetailsService: factory.NewTorrentDetailsService(factory.Mock)}.Start()
}
