package API

import (
	util "github.com/04Akaps/go-util/log"
	"github.com/04Akaps/metting/aws"
	"github.com/04Akaps/metting/config"
	"github.com/04Akaps/metting/db/HDD"
	"github.com/04Akaps/metting/db/SSD"
	"github.com/04Akaps/metting/module/API/network"
	"github.com/04Akaps/metting/module/API/service"
)

type API struct {
	cfg *config.Config
	log *util.Log

	network *network.Network
	service service.ServiceImpl
}

func NewAPI(
	cfg *config.Config,
	log *util.Log,
	db *HDD.DB,
	redis SSD.Redis,
	aws *aws.GoAWS,
) *API {
	a := &API{cfg: cfg, log: log}

	a.service = service.NewService(cfg, log, db, aws, redis)
	a.network = network.NewNetwork(cfg, log, a.service)

	go func() {
		a.network.StartServer()
	}()

	return a
}
