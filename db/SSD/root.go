package SSD

import (
	"encoding/json"
	"github.com/04Akaps/metting/config"
	. "github.com/go-redis/redis/v7"
	"log"
	"time"
)

type redis struct {
	config *config.Config
	client *Client
}

type Redis interface {
	Save(key string, data interface{}, time time.Duration) error
	Get(key string) (interface{}, error)
	GetAllKeys(pattern string) ([]string, error)
}

func NewRedis(config *config.Config) Redis {
	r := &redis{config: config}

	r.client = NewClient(&Options{
		Addr:     config.Redis.Url,
		Password: config.Redis.Password,
		Username: config.Redis.User,
		DB:       config.Redis.DB,
	})

	if _, err := r.client.Ping().Result(); err != nil {
		panic(err)
	} else {
		log.Println("Success To Connect Redis")
		return r
	}
}

func (r *redis) Save(key string, data interface{}, time time.Duration) error {
	_, err := r.client.Set(key, data, time).Result()
	return err
}

func (r *redis) Get(key string) (interface{}, error) {
	if d, err := r.client.Get(key).Bytes(); err != nil {
		return nil, err
	} else {
		var res interface{}

		if err = json.Unmarshal(d, &res); err != nil {
			return nil, err
		} else {
			return res, nil
		}
	}
}

func (r *redis) GetAllKeys(pattern string) ([]string, error) {
	return r.client.Keys(pattern).Result()
}
