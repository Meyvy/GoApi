package db

import (
	"context"
	"fmt"
	"monitor/model"
	"time"
)

func AddUser(s model.RegisterUserRequest) (int64, error) {
	ctx := context.Background()
	id := Rdb.Incr(ctx, "user_id")
	if id.Err() != nil {
		return 0, id.Err()
	}
	key := fmt.Sprintf("user:%d", id.Val())
	Rdb.HSet(ctx, key, model.User{
		UserName:  *s.UserName,
		PassWord:  *s.PassWord,
		CreatedAt: time.Now().Format(time.ANSIC),
	})
	return id.Val(), nil
}

func GetUser(id int64) (model.User, error) {
	key := fmt.Sprintf("user:%d", id)
	ctx := context.Background()
	u := model.User{}
	res := Rdb.HGetAll(ctx, key)
	if res.Err() != nil {
		return u, res.Err()
	}
	u.UserName = res.Val()["user_name"]
	u.PassWord = res.Val()["password"]
	u.CreatedAt = res.Val()["created_at"]
	return u, nil
}
