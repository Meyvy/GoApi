package app

import (
	"context"
	"monitor/config"
	"monitor/db"
	"monitor/model"
	"net/http"
	"strconv"
	"time"
)

func sendRequest(l model.Link, userId int64) model.Request {
	result := model.Request{}
	cli := &http.Client{
		Timeout: config.ClinetTimeOut,
	}
	req, _ := http.NewRequestWithContext(context.Background(),
		l.Method, l.Url, nil)
	response, _ := cli.Do(req)
	result.CreatedAt = time.Now().Format(time.ANSIC)
	result.Status = "failed"
	if response != nil && response.StatusCode < 300 && response.StatusCode >= 200 {
		result.Status = "success"
	}
	db.AddRequest(result, userId, l.LinkID)
	if result.Status == "failed" {
		db.IncreaseFailure(userId, l.LinkID)
	}
	return result
}

func updateUser(userID int64) {
	links, err := db.GetAllLink(userID)
	if err != nil || len(links) == 0 {
		return
	}
	for _, v := range links {
		sendRequest(v, userID)
	}
}
func update() {
	ctx := context.Background()
	user_id := db.Rdb.Get(ctx, "user_id")
	if user_id.Err() != nil {
		for {
			user_id = db.Rdb.Get(ctx, "user_id")
			if user_id.Err() == nil {
				break
			}
			<-time.After(time.Second * 1)
		}
	}
	n, _ := strconv.ParseInt(user_id.Val(), 10, 64)
	for i := int64(1); i <= n; i++ {
		go updateUser(i)
	}

}

func monitor() {
	for {
		update()
		<-time.After(config.WaitDuration)
	}
}
