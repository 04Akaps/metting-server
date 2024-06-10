package main

import (
	"flag"
	util "github.com/04Akaps/go-util/log"
	"github.com/04Akaps/metting/aws"
	"github.com/04Akaps/metting/config"
	"github.com/04Akaps/metting/db/HDD"
	"github.com/04Akaps/metting/db/SSD"
	"github.com/04Akaps/metting/module/API"
	"github.com/04Akaps/metting/module/Internal"
	"go.uber.org/fx"
)

var cfgPath = flag.String("cfg", "./config.toml", "for config")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*cfgPath)
	log := util.SetLog(cfg.Info.Log)

	fx.New(
		fx.Provide(func() *config.Config { return cfg }),
		fx.Provide(func() *util.Log { return log }),
		fx.Provide(func() *HDD.DB { return HDD.NewDB(cfg, log) }),
		fx.Provide(func() SSD.Redis { return SSD.NewRedis(cfg) }),
		fx.Provide(func() *aws.GoAWS { return aws.NewAWS(cfg, log) }),

		fx.Provide(Internal.NewInternal, API.NewAPI),
		fx.Invoke(func(*Internal.Internal) {}, func(*API.API) {}),
	).Run()
}
