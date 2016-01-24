package library

import (
	"config"
	"github.com/garyburd/redigo/redis"
	"logger"
)

type RedisConnectionPool struct {
	redis_pool *redis.Pool
}

func NewRedisConnectionPool() *RedisConnectionPool {
	return &RedisConnectionPool{}
}

func (r *RedisConnectionPool) newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Redis_Address)
			if err != nil {
				logger.Get().Debug("Error while connecting : " + err.Error())
			}
			return c, err
		},
	}

}

func (r *RedisConnectionPool) GetConnection() redis.Conn {
	if r.redis_pool == nil {
		r.redis_pool = r.newPool()
	}
	return r.redis_pool.Get()
}
