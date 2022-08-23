package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/timhuynh94/TargetChallenge/models"
	"log"
)

// RDBService redis client service
type RDBService struct {
	client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	ctx    = context.TODO()
)

func NewRdbClient() *RDBService {
	// Connect to REDIS DB for local storage of product details
	rdb := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}
	return &RDBService{client: rdb}
}

func (rdb *RDBService) getProductFromRedis(id string) (*models.RespBody, error) {
	var v models.RespBody
	val, err := rdb.client.Get(ctx, id).Result()

	if err != nil {
		return nil, ErrNil
	}
	unmarshalErr := json.Unmarshal([]byte(val), &v)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return &v, nil
}

func (rdb *RDBService) setProductToRedis(id string, val models.RespBody) error {
	v, marshalErr := json.Marshal(val)
	if marshalErr != nil {
		return marshalErr
	}
	err := rdb.client.Set(ctx, id, v, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
