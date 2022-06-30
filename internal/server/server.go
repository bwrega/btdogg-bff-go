package server

import "fmt"
import "net/http"
import "github.com/gin-gonic/gin"

type HttpServer struct {
	PortNumber int
}

func (s HttpServer) Start() {
	fmt.Println("Starting server, lol")
	router := gin.Default()
	router.GET("/hello", hello)
	router.Run(fmt.Sprint(":", s.PortNumber))
}

func hello(c *gin.Context) {
	c.String(http.StatusTeapot, "Lol I'm a teapot!")
}
