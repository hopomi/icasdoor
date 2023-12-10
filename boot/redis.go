package boot

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/redis/go-redis/v9"
)

var DefaultRedisClient = redis.NewClient(&redis.Options{
	Addr:     config.DefaultString("redisconn.default.addr", ""),
	Password: config.DefaultString("redisconn.default.passwd", ""),
	DB:       1,
})
