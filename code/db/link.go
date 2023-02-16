package db

import (
	"context"
	"fmt"
	Err "monitor/error"
	"monitor/model"
	"strconv"
	"time"
)

// Insert a verified link request object  into the database
// our key for the link is ["user:<user_id>:link<link_id>"]
func AddLink(l model.RegisterLinkRequest, id int64) (int64, error) {
	numberLinkKey := fmt.Sprintf("user:%d:link", id)
	ctx := context.Background()
	linkId := Rdb.Incr(ctx, numberLinkKey)
	if linkId.Err() != nil {
		return 0, linkId.Err()
	}
	if linkId.Val() > 20 {
		for {
			res := Rdb.Decr(ctx, numberLinkKey)
			if res.Err() == nil {
				break
			}
			<-time.After(time.Second * 1)
		}
		return 0, Err.ErrMaxNumLink
	}
	key := fmt.Sprintf("user:%d:link:%d", id, linkId.Val())
	Rdb.HSet(ctx, key, model.Link{
		LinkID:     linkId.Val(),
		Url:        *l.Url,
		ThreshHold: *l.ThreshHold,
		CreatedAt:  time.Now().Format(time.ANSIC),
		Failures:   0,
		Method:     *l.Method,
	})
	return linkId.Val(), nil
}

// gets a  link with the user and link id
func GetLink(userID int64, linkID int64) (model.Link, error) {
	key := fmt.Sprintf("user:%d:link:%d", userID, linkID)
	ctx := context.Background()
	l := model.Link{}
	res := Rdb.HGetAll(ctx, key)
	if res.Err() != nil {
		return l, res.Err()
	}
	l.LinkID = linkID
	l.Failures, _ = strconv.Atoi(res.Val()["failures"])
	l.CreatedAt = res.Val()["created_at"]
	l.Method = res.Val()["method"]
	l.Url = res.Val()["url"]
	l.ThreshHold, _ = strconv.Atoi(res.Val()["thresh_hold"])
	return l, nil
}

// gets all the links of a user with its id
func GetAllLink(userId int64) ([]model.Link, error) {
	list := make([]model.Link, 0)
	numberLinkKey := fmt.Sprintf("user:%d:link", userId)
	ctx := context.Background()
	linkId := Rdb.Get(ctx, numberLinkKey)
	if linkId.Err() != nil {
		return nil, linkId.Err()
	}
	n, _ := strconv.ParseInt(linkId.Val(), 10, 64)
	for i := int64(1); i <= n; i++ {
		l, err := GetLink(userId, i)
		if err == nil {
			list = append(list, l)
		}
	}
	return list, nil
}

func IncreaseFailure(userId, link_id int64) {
	key := fmt.Sprintf("user:%d:link:%d", userId, link_id)
	ctx := context.Background()
	res := Rdb.HIncrBy(ctx, key, "failures", 1)
	if res.Err() != nil {
		for {
			res = Rdb.HIncrBy(ctx, key, "failures", 1)
			if res.Err() == nil {
				break
			}
			<-time.After(time.Second * 1)
		}
	}
}
