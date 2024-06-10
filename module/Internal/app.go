package Internal

import (
	util "github.com/04Akaps/go-util/log"
	"github.com/04Akaps/metting/config"
	"github.com/04Akaps/metting/db/HDD"
	"github.com/04Akaps/metting/db/SSD"
	"github.com/04Akaps/metting/module/Internal/remove"
	"github.com/robfig/cron"
	"os"
	"os/signal"
)

type Internal struct {
	cfg   *config.Config
	log   *util.Log
	db    *HDD.DB
	redis SSD.Redis
	c     *cron.Cron
}

func NewInternal(
	cfg *config.Config,
	log *util.Log,
	db *HDD.DB,
	redis SSD.Redis,
) *Internal {
	i := &Internal{
		cfg:   cfg,
		log:   log,
		db:    db,
		redis: redis,
		c:     cron.New(),
	}

	i.c.Start()

	go remove.NewFileRemover(cfg, log, db, redis, i.c)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		i.c.Stop()
	}()

	return i
}
