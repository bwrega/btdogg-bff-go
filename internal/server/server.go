package server

import (
	"fmt"
	"github.com/bwrega/btdogg-bff-go/internal/details"
	"github.com/bwrega/btdogg-bff-go/internal/search"
)
import "net/http"
import "github.com/gin-gonic/gin"

type HttpServer struct {
	PortNumber int
	DetailsService details.TorrentDetailsService
	SearchService search.TorrentSearchService
}

func (s HttpServer) Start() {
	fmt.Println("Starting server, lol")
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.GET("/details/:torrentHash", s.getDetails)
	v1.GET("/search", s.search)
	router.Run(fmt.Sprint(":", s.PortNumber))
}

func (s HttpServer) getDetails(c *gin.Context) {
	torrentHash := details.TorrentHash(c.Param("torrentHash"))
	torrentDetails := s.DetailsService.GetDetails(torrentHash)
	if torrentDetails == nil {
		c.String(http.StatusNotFound, "Torrent with id " + string(torrentHash) + " not found")
	} else {
		c.JSON(http.StatusOK, torrentDetails)
	}
}

func (s HttpServer) search(c *gin.Context) {
	query := c.DefaultQuery("q", "")
	c.JSON(http.StatusOK, s.SearchService.Search(query))
}
