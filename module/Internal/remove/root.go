package remove

import (
	"fmt"
	util "github.com/04Akaps/go-util/log"
	"github.com/04Akaps/metting/config"
	"github.com/04Akaps/metting/db/HDD"
	"github.com/04Akaps/metting/db/SSD"
	"github.com/04Akaps/metting/types"
	"github.com/robfig/cron"
)

type FileRemover struct {
	cfg *config.Config
	log *util.Log

	redis SSD.Redis
	db    *HDD.DB
	c     *cron.Cron
}

func NewFileRemover(
	cfg *config.Config,
	log *util.Log,
	db *HDD.DB,
	redis SSD.Redis,
	c *cron.Cron,
) {
	fe := FileRemover{
		cfg:   cfg,
		log:   log,
		redis: redis,
		db:    db,
		c:     c,
	}

	go fe.fileRemove()
}

func (fe *FileRemover) fileRemove() {
	imgPath := types.IMG_PATH
	fmt.Println(imgPath)
}
