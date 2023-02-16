package config

import (
	"time"

	"github.com/redis/go-redis/v9"
)

// our database configs it should be loaded from a file or env variable
const addr string = "redis:6379"
const password = ""
const db int = 0

// our redis database option
var DbOptions = redis.Options{
	Addr:     addr,
	Password: password,
	DB:       db,
}

// our api reference for people to refer to
const Help = "anything"

// link config
const MaxNumLink = 20

// fetch config
const ClinetTimeOut = time.Second * 30

// server config
const Addr = ":8080"
const ReadTimeOut = time.Second * 30
const WriteTimeOut = time.Second * 30
const IdleTimeOut = time.Second * 30

//update config

const WaitDuration = time.Minute * 10

const TokenDuration = time.Hour * 2
