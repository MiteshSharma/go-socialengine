package library

import(
	"github.com/garyburd/redigo/redis"
	"logger"
)

type RedisHash struct {
	redis_connection redis.Conn
	hash_name string
	page_size int
}

var redisHash map[string]*RedisHash

const (
	prefixhash = "hash-"
)

func GetRedisHashInstance(hashName string) *RedisHash {
	if redisHash == nil {
		redisHash = make(map[string]*RedisHash)
	}
	if redisHash[hashName] == nil {
		hashInstance := newRedisHash()
		redisHash[hashName] = hashInstance
		redisHash[hashName].hash_name = prefixhash+hashName
	}
	return redisHash[hashName]
}

func newRedisHash() *RedisHash {
	return &RedisHash{}
}

func (r *RedisHash) GetRedisConnection() redis.Conn {
	if r.redis_connection == nil {
		r.redis_connection = NewRedisConnectionPool().GetConnection()
	}
	return r.redis_connection;
}

func (r *RedisHash) CloseConn() {
	redisConn := r.GetRedisConnection()
	if redisConn != nil {
		redisConn.Close()
	}
}

func (r *RedisHash) AddHashValue(key string, value string) {
	if r.redis_connection != nil {
		_,err := r.redis_connection.Do("HSET", r.hash_name, key, value)
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
	}
}

func (r *RedisHash) RemoveHashValue(key string) {
	if r.redis_connection != nil {
		_,err := r.redis_connection.Do("HDEL", r.hash_name, key)
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
	}
}

func (r *RedisHash) TotalValue() int {
	if r.redis_connection != nil {
		count,err := redis.Int(r.redis_connection.Do("HGETALL", r.hash_name))
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
		return count
	}
	return 0
}

func (r *RedisHash) GetHashValue(key string) string {
	if r.redis_connection != nil {
		value,err := redis.String(r.redis_connection.Do("HGET", r.hash_name, key))
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
		return value
	}
	return ""
}

func (r *RedisHash) SetHashFromMap(mapValues map[string]string) {
	if r.redis_connection != nil {
		r.redis_connection.Send("MULTI")
		for key := range mapValues {
			r.redis_connection.Send("HSET", r.hash_name, key, mapValues[key])
		}
		_, err := r.redis_connection.Do("EXEC")
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
	}
}