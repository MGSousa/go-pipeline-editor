package cache

import (
	rd "github.com/go-redis/redis"
)

// Redis kv cache registry
type Redis struct {
	client *rd.Client

	Address string
	DB      int
}

// connect to a new Redis instance
func (r *Redis) connect() error {
	if r.client == nil {
		r.client = rd.NewClient(&rd.Options{
			Addr: r.Address,
			DB:   r.DB,
		})
		// ping application
		if _, err := r.client.Ping().Result(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Redis) set(key string, value interface{}) (string, error) {
	return r.client.Set(key, value, 0).Result()
}

func (r *Redis) get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *Redis) del(key string) (int64, error) {
	return r.client.Del(key).Result()
}
