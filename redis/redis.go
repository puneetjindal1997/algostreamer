package redis

import (
	"os"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type redisDatabase struct {
	client *goredislib.Client
}

var RedSync *redsync.Redsync

/*
 *	Function to create redis database
 *	return redis client
 */
func createRedisDatabase() (Database, error) {
	client := goredislib.NewClient(&goredislib.Options{
		// Addr: constant.RedisHost + ":" + constant.RedisPort,
		Addr:     os.Getenv("RedisHost") + ":" + os.Getenv("RedisPort"),
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, &CreateDatabaseError{}
	}
	pool := goredis.NewPool(client)
	RedSync = redsync.New(pool)
	return &redisDatabase{client: client}, nil
}

/*
 *	Function to set the value to a perticular key in redis
 *	retrn string, error
 */
func (r *redisDatabase) Set(key string, value interface{}) (string, error) {
	_, err := r.client.Set(r.client.Context(), key, value, 300000000000000000).Result()
	if err != nil {
		return generateError("set", err)
	}
	return key, nil
}

/*
 *	Function to get the value for a perticular key in redis
 *	return interface{}, error
 */
func (r *redisDatabase) Get(key string) (interface{}, error) {
	value, err := r.client.Get(r.client.Context(), key).Result()
	if err != nil {
		return generateError("get", err)
	}
	return value, nil
}

/*
 *	Function to delete the perticular key and its data from redis
 *	return string, error
 */
func (r *redisDatabase) Delete(key string) (string, error) {
	_, err := r.client.Del(r.client.Context(), key).Result()
	if err != nil {
		return generateError("delete", err)
	}
	return key, nil
}

/*
 *	Function to hash set for resdis
 *	return error
 */
func (r *redisDatabase) HSet(key, field string, data interface{}) (err error) {
	_, err = r.client.HSet(r.client.Context(), key, field, data).Result()
	return err
}

/*
 *	Function to hash get for redis
 *	return interface{}, error
 */
func (r *redisDatabase) HGet(key, field string) (data interface{}, err error) {
	data, err = r.client.HGet(r.client.Context(), key, field).Result()
	return data, err
}

/*
 *	Function to generate error if occure in redis
 */
func generateError(operation string, err error) (string, error) {
	if err == goredislib.Nil {
		return "", &OprationError{operation}
	}
	return "", &DownError{}
}
