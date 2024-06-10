package network

import (
	util "github.com/04Akaps/go-util/log"
	"github.com/04Akaps/metting/config"
	"github.com/04Akaps/metting/module/API/service"
	"github.com/gin-gonic/gin"
)

type Network struct {
	service service.ServiceImpl

	engine *gin.Engine

	port string
	cfg  *config.Config
	log  *util.Log
}

func NewNetwork(
	cfg *config.Config,
	log *util.Log,
	service service.ServiceImpl,
) *Network {
	n := &Network{
		service: service,
		engine:  gin.New(),
		cfg:     cfg,
		log:     log,
		port:    cfg.Info.Port,
	}

	n.setGin()

	userRouter(n)

	return n
}

func (n *Network) StartServer() error {
	n.log.InfoLog("Start API Port", "port", n.cfg.Info.Port, "env", n.cfg.Info.Service)
	return n.engine.Run(n.port)
}
