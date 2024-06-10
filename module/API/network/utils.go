package network

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
)

type Router int8

const (
	GET Router = iota
	POST
	DELETE
	PUT
)

type header struct {
	Result int    `json:"result"`
	Data   string `json:"data"`
}

type response struct {
	*header
	Result interface{} `json:"result"`
}

func res(c *gin.Context, s int, res interface{}, data ...string) {
	c.JSON(s, &response{
		header: &header{Result: s, Data: strings.Join(data, ",")},
		Result: res,
	})
}

func (n *Network) setGin() {
	n.engine.Use(gin.Logger())
	n.engine.Use(gin.Recovery())
	n.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		ExposeHeaders:    []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))
}

func (n *Network) Router(r Router, path string, handler gin.HandlerFunc) {
	switch r {
	case GET:
		n.engine.GET(path, handler)
	case POST:
		n.engine.POST(path, handler)
	case PUT:
		n.engine.DELETE(path, handler)
	case DELETE:
		n.engine.PUT(path, handler)
	default:
		n.log.CritLog("Failed To Register API", "path", path)
	}
}
