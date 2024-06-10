package HDD

import (
	"database/sql"
	util "github.com/04Akaps/go-util/log"
	"github.com/04Akaps/metting/config"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	cfg *config.Config
	log *util.Log

	db *sql.DB
}

func NewDB(cfg *config.Config, log *util.Log) *DB {
	d := &DB{cfg: cfg, log: log}

	var err error

	if d.db, err = sql.Open(cfg.DB.Database, cfg.DB.URL); err != nil {
		panic(err)
	} else if err = d.db.Ping(); err != nil {
		panic(err)
	} else {
		return d
	}

}
