package model

type ErrResponse struct {
	Status    string `json:"status"`
	ErrorMsg  string `json:"error_msg"`
	Help      string `json:"help"`
	ErrorCode int    `json:"-"`
}

type RegisterUserResponse struct {
	Status string `json:"status"`
	UserId int64  `json:"user_id"`
}

type RegisterLinkResponse struct {
	Status string `json:"status"`
	LinkId int64  `json:"link_id"`
}

type TokenResponse struct {
	Status  string `json:"status"`
	ExpDate string `json:"exp_date"`
	Token   string `json:"token"`
}

type LinkResponse struct {
	Link
	Requests []Request `json:"requests"`
}
