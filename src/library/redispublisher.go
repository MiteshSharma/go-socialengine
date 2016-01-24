package library

import (
	"github.com/garyburd/redigo/redis"
	"logger"
)

type RedisPublisher struct {
	channel          string
	redis_connection redis.Conn
}

var publisher *RedisPublisher

func GetRedisPublisherInstance(channel string) *RedisPublisher {
	if publisher == nil {
		publisher := newRedisPublisher()
		publisher.channel = channel
		publisher.GetRedisConnection()
	}
	return publisher
}

func newRedisPublisher() *RedisPublisher {
	return &RedisPublisher{}
}

func (r *RedisPublisher) GetRedisConnection() redis.Conn {
	if r.redis_connection == nil {
		r.redis_connection = NewRedisConnectionPool().GetConnection()
	}
	return r.redis_connection
}

func (r *RedisPublisher) CloseConn() {
	redisConn := r.GetRedisConnection()
	if redisConn != nil {
		redisConn.Close()
	}
}

func (r *RedisPublisher) PublishChannelMessage(channel string, message string) {
	if r.redis_connection != nil {
		_, err := r.redis_connection.Do("PUBLISH", channel, message)
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
	}
}

func (r *RedisPublisher) PublishMessage(message string) {
	if r.redis_connection != nil {
		_, err := r.redis_connection.Do("PUBLISH", r.channel, message)
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
	}
}
