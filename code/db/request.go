package db

import (
	"context"
	"fmt"
	"monitor/model"
	"strconv"
	"time"
)

func AddRequest(req model.Request, userId, LinkId int64) error {
	now := time.Now()
	current_time := fmt.Sprintf("%d:%d:%d", now.Year(), int(now.UTC().Month()), now.Day())
	requestIdKey := fmt.Sprintf("user:%d:link:%d:%s", userId, LinkId, current_time)
	ctx := context.Background()
	requestId := Rdb.Incr(ctx, requestIdKey)
	if requestId.Err() != nil {
		return requestId.Err()
	}
	requestKey := fmt.Sprintf("user:%d:link:%d:%s:%d", userId, LinkId, current_time, requestId.Val())
	Rdb.HSet(ctx, requestKey, req)
	return nil
}

func GetRequest(key string) (model.Request, error) {
	ctx := context.Background()
	r := model.Request{}
	res := Rdb.HGetAll(ctx, key)
	if res.Err() != nil {
		return r, res.Err()
	}
	r.Status = res.Val()["status"]
	r.CreatedAt = res.Val()["created_at"]
	return r, nil
}

func GetTodayRequest(userId, LinkId int64) ([]model.Request, error) {
	now := time.Now()
	current_time := fmt.Sprintf("%d:%d:%d", now.Year(), int(now.UTC().Month()), now.Day())
	requestIdKey := fmt.Sprintf("user:%d:link:%d:%s", userId, LinkId, current_time)
	var requestKey string
	ctx := context.Background()
	requestCount := Rdb.Get(ctx, requestIdKey)
	if requestCount.Err() != nil {
		return nil, requestCount.Err()
	}
	n, _ := strconv.ParseInt(requestCount.Val(), 10, 64)
	list := make([]model.Request, 0)
	for i := int64(1); i <= n; i++ {
		requestKey = fmt.Sprintf("user:%d:link:%d:%s:%d", userId, LinkId, current_time, i)
		request, err := GetRequest(requestKey)
		if err == nil {
			list = append(list, request)
		}
	}
	return list, nil
}
