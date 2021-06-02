package redisstore

import (
	"github.com/gomodule/redigo/redis"
)

const keyModelVersion = "modelVersion"

var emptyItem = struct {
	RedisNil bool `redis:"redisNil"`
}{true}

// redisStore defines base struct for Redis stores.
type redisStore struct {
	conn         redis.Conn
	baseKey      string
	modelVersion int
}

func (store *redisStore) DoEXISTS(key string) (bool, error) {
	return redis.Bool(store.conn.Do("EXISTS", key))
}

func (store *redisStore) DoGETBool(key string) (bool, error) {
	return redis.Bool(store.conn.Do("GET", key))
}

func (store *redisStore) DoGETInt(key string) (int, error) {
	return redis.Int(store.conn.Do("GET", key))
}

func (store *redisStore) DoGETInt64(key string) (int64, error) {
	return redis.Int64(store.conn.Do("GET", key))
}

func (store *redisStore) DoGETString(key string) (string, error) {
	return redis.String(store.conn.Do("GET", key))
}

func (store *redisStore) DoSET(key string, value interface{}, ttl int) error {
	if _, err := store.conn.Do("SET", key, value); err != nil {
		return err
	}
	if ttl > 0 {
		store.conn.Do("EXPIRE", key, ttl)
	}
	return nil
}

func (store *redisStore) DoHGETALL(key string, dest interface{}) error {
	v, err := redis.Values(store.conn.Do("HGETALL", key))
	if err != nil {
		return err
	} else if len(v) == 0 {
		return redis.ErrNil
	}
	return redis.ScanStruct(v, dest)
}

func (store *redisStore) DoHMSET(key, v interface{}, ttl int) error {
	if _, err := store.conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(v)...); err != nil {
		return err
	}
	if ttl > 0 {
		store.conn.Do("EXPIRE", key, ttl)
	}
	return nil
}

func (store *redisStore) DoLRANGEInts(key string) ([]int, error) {
	return redis.Ints(store.conn.Do("LRANGE", key, 0, -1))
}

func (store *redisStore) DoLRANGEInt64s(key string) ([]int64, error) {
	return redis.Int64s(store.conn.Do("LRANGE", key, 0, -1))
}

func (store *redisStore) DoLRANGEStrings(key string) ([]string, error) {
	return redis.Strings(store.conn.Do("LRANGE", key, 0, -1))
}

func (store *redisStore) DoRPUSH(key string, items ...interface{}) error {
	if _, err := store.conn.Do("RPUSH", redis.Args{}.Add(key).Add(items...)...); err != nil {
		return err
	}
	return nil
}

func (store *redisStore) DoEXPIRE(key string, ttl int) error {
	_, err := store.conn.Do("EXPIRE", key, ttl)
	return err
}

func (store *redisStore) DoDEL(keys ...string) (count int, err error) {
	switch len(keys) {
	case 0:
		return 0, nil
	case 1:
		count, err = redis.Int(store.conn.Do("DEL", keys[0]))
	default:
		args := make([]interface{}, len(keys))
		for i, k := range keys {
			args[i] = k
		}
		count, err = redis.Int(store.conn.Do("DEL", args...))
	}
	if err != nil && err == redis.ErrNil {
		err = nil
	}
	return
}
