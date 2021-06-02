package redis

import (
	// See https://godoc.org/github.com/gomodule/redigo/redis.
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/gomodule/redigo/redis"
)

// config defines the Redis connection configurations.
type config struct {
	Host           string
	Port           string
	Database       string
	Password       string
	MaxConnections string
}

// redisStore defines base struct for Redis stores.
type redisStore struct {
	conn    redis.Conn
	baseKey string
}

var redisPool *redis.Pool

// Init initializes Redis connection pool using the specified configurations. Do NOT forget to defer closing the connection pool.
func Init() {
	logger.Println("redis", "Initializing Redis connection pool...")

	// Load environment variables.
	cfg := config{
		Host:           os.Getenv(envvar.RedisHost),
		Port:           os.Getenv(envvar.RedisPort),
		Database:       os.Getenv(envvar.RedisDatabase),
		Password:       os.Getenv(envvar.RedisPassword),
		MaxConnections: os.Getenv(envvar.RedisMaxConnections),
	}
	logger.Println("redis", fmt.Sprintf(`Host = %s, Port = %s, Database = %s, Password = %v, MaxConnections = %v`,
		cfg.Host, cfg.Port, cfg.Database, cfg.Password != "", cfg.MaxConnections))

	// Create Redis connection pool.
	var err error
	redisPool, err = newPool(cfg)
	if err != nil {
		logger.Println("redis", fmt.Sprintf("ERROR: Init: %v", err))
	}

	conn := GetConnection()
	defer conn.Close()

	// Save init data.
	init := struct {
		Env  string `redis:"env"`
		Time string `redis:"time"`
	}{os.Getenv(envvar.Environment), time.Now().String()}
	if _, err = conn.Do("HMSET", redis.Args{}.Add("init").AddFlat(&init)...); err != nil {
		logger.Println("redis", fmt.Sprintf("ERROR: Init: %v", err))
	}
}

func newPool(cfg config) (*redis.Pool, error) {
	var maxConns int
	if cfg.Host == "" {
		return nil, errors.New("newPool: Host is undefined")
	} else if cfg.Port == "" {
		return nil, errors.New("newPool: Port is undefined")
	}

	dialOptions := make([]redis.DialOption, 0, 2)
	if cfg.Database != "" {
		i, err := helper.StringToInt(cfg.Database)
		if err != nil {
			return nil, errors.New("newPool: Database must be integer")
		}
		dialOptions = append(dialOptions, redis.DialDatabase(i))
	}
	if cfg.Password != "" {
		dialOptions = append(dialOptions, redis.DialPassword(cfg.Password))
	}

	if cfg.MaxConnections != "" {
		i, err := helper.StringToInt(cfg.MaxConnections)
		if err != nil {
			return nil, errors.New("newPool: MaxConnections must be integer")
		}
		maxConns = i
	}

	redisAddr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisAddr, dialOptions...)
			if err != nil {
				return nil, fmt.Errorf("redis.Dial: %v", err)
			}
			return conn, err
		},
		MaxIdle:     maxConns,
		IdleTimeout: 5 * time.Minute,
	}
	return pool, nil
}

// Close closes the Redis connection pool.
func Close() {
	redisPool.Close()
}

// GetConnection returns a Redis connection. Do NOT forget to defer closing the connection.
func GetConnection() redis.Conn {
	return redisPool.Get()
}
