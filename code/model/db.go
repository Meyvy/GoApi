package model

type User struct {
	UserName  string `redis:"user_name"`
	PassWord  string `redis:"password"`
	CreatedAt string `redis:"created_at"`
}

type Link struct {
	LinkID     int64  `redis:"link_id"`
	Url        string `redis:"url"`
	ThreshHold int    `redis:"thresh_hold"`
	CreatedAt  string `redis:"created_at"`
	Method     string `redis:"method"`
	Failures   int    `redis:"failures"`
}

type Request struct {
	Status    string `redis:"status"`
	CreatedAt string `redis:"created_at"`
}
