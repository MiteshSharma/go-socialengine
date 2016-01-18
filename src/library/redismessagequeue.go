package library

import(
	"github.com/garyburd/redigo/redis"
	"logger"
)

type RedisMessageQueue struct {
	redis_connection redis.Conn
	list_name string
}

var redisMessageQueue *RedisMessageQueue

func GetRedisMessageQueueInstance() *RedisMessageQueue {
	if redisMessageQueue == nil {
		redisMessageQueue = newRedisMessageQueue()
	}
	return redisMessageQueue
}

func newRedisMessageQueue() *RedisMessageQueue {
	return &RedisMessageQueue{}
}

func (r *RedisMessageQueue) GetRedisConnection() redis.Conn {
	if r.redis_connection == nil {
		r.redis_connection = NewRedisConnectionPool().GetConnection()
	}
	return r.redis_connection;
}

func (r *RedisMessageQueue) CloseConn() {
	redisConn := r.GetRedisConnection()
	if redisConn != nil {
		redisConn.Close()
	}
}

func (r *RedisMessageQueue) PostMessage(list, message string) {
	if r.redis_connection != nil {
		_,err := redis.String(r.redis_connection.Do("RPUSH", list, message))
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
	}
}

func PostMessageOnQueue(queue_name, message string) {
	redisMessageQueue := GetRedisMessageQueueInstance()
	_,err := redisMessageQueue.GetRedisConnection().Do("RPUSH", queue_name, message)
	if err != nil {
		logger.Get().Debug("Error is : " + err.Error())
	}
	redisMessageQueue.CloseConn()
}