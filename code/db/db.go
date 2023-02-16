package db

import (
	"monitor/config"

	"github.com/redis/go-redis/v9"
)

// our db connection there is only one since redis is thread safe
var Rdb = redis.NewClient(&config.DbOptions)
