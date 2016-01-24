package library

import (
	"github.com/garyburd/redigo/redis"
	"logger"
)

type RedisLeaderboard struct {
	redis_connection redis.Conn
	leaderboard_name string
	page_size        int
}

var redisLeaderboard map[string]*RedisLeaderboard

const (
	prefixleaderboard = "lb-"
)

func GetRedisLeaderBoardInstance(leaderboardName string) *RedisLeaderboard {
	if redisLeaderboard == nil {
		redisLeaderboard = make(map[string]*RedisLeaderboard)
	}
	if redisLeaderboard[leaderboardName] == nil {
		leaderboardInstance := newRedisLeaderboard()
		redisLeaderboard[leaderboardName] = leaderboardInstance
		redisLeaderboard[leaderboardName].leaderboard_name = prefixleaderboard + leaderboardName
	}
	return redisLeaderboard[leaderboardName]
}

func newRedisLeaderboard() *RedisLeaderboard {
	return &RedisLeaderboard{}
}

func (r *RedisLeaderboard) GetRedisConnection() redis.Conn {
	if r.redis_connection == nil {
		r.redis_connection = NewRedisConnectionPool().GetConnection()
	}
	return r.redis_connection
}

func (r *RedisLeaderboard) CloseConn() {
	redisConn := r.GetRedisConnection()
	if redisConn != nil {
		redisConn.Close()
	}
}

func (r *RedisLeaderboard) GetLeaderboardName() string {
	return r.leaderboard_name
}

func (r *RedisLeaderboard) setPageSize(pageSize int) {
	if pageSize < 1 {
		r.page_size = 20
	} else {
		r.page_size = pageSize
	}
}

func (r *RedisLeaderboard) AddMember(member string, score float64) {
	if r.redis_connection != nil {
		_, err := r.redis_connection.Do("ZADD", r.leaderboard_name, score, member)
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
	}
}

func (r *RedisLeaderboard) RemoveMember(member string) {
	if r.redis_connection != nil {
		_, err := r.redis_connection.Do("ZREM", r.leaderboard_name, member)
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
	}
}

func (r *RedisLeaderboard) TotalMembers() int {
	if r.redis_connection != nil {
		count, err := redis.Int(r.redis_connection.Do("ZCARD", r.leaderboard_name))
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
		return count
	}
	return 0
}

func (r *RedisLeaderboard) MemberRank(member string) int {
	if r.redis_connection != nil {
		count, err := redis.Int(r.redis_connection.Do("ZREVRANK", r.leaderboard_name, member))
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
		return count
	}
	return -1
}

func (r *RedisLeaderboard) MemberScore(member string) int {
	if r.redis_connection != nil {
		count, err := redis.Int(r.redis_connection.Do("ZSCORE", r.leaderboard_name, member))
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
		return count
	}
	return 0
}

func (r *RedisLeaderboard) GetLeaders(currentPage int) []string {
	if r.redis_connection != nil {
		if currentPage < 1 {
			currentPage = 1
		}
		startOffset := (currentPage - 1) * r.page_size
		if startOffset < 0 {
			startOffset = 0
		}
		endOffset := startOffset + r.page_size - 1
		leaderData, err := redis.Strings(r.redis_connection.Do("ZREVRANGE", r.leaderboard_name, startOffset, endOffset))
		if err != nil {
			logger.Get().Debug("Error is : " + err.Error())
		}
		return leaderData
	}
	return nil
}
