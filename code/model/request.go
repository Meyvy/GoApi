package model

type TokenRequest struct {
	UserId   *int64  `json:"user_id"`
	PassWord *string `json:"password"`
}

type RegisterUserRequest struct {
	UserName *string `json:"user_name"`
	PassWord *string `json:"password"`
}

type RegisterLinkRequest struct {
	Url       *string `json:"url"`
	Threshold *int    `json:"threshold"`
	Method    *string `json:"method"`
}

type LinkRequest struct {
	LinkId *int64 `json:"link_id"`
}
