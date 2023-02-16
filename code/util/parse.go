package util

import (
	"encoding/json"
	Err "monitor/error"
	"monitor/model"
	"net/http"
)

func ParseUserSignUp(r *http.Request) (model.RegisterUserRequest, error) {
	s := model.RegisterUserRequest{}
	json.NewDecoder(r.Body).Decode(&s)
	if s.PassWord == nil || s.UserName == nil {
		return s, Err.ErrMissingFields
	}
	return s, nil
}

func ParseToken(r *http.Request) (model.TokenRequest, error) {
	s := model.TokenRequest{}
	json.NewDecoder(r.Body).Decode(&s)
	if s.UserId == nil || s.PassWord == nil {
		return s, Err.ErrMissingFields
	}
	return s, nil
}

func ParseRegisterLink(r *http.Request) (model.RegisterLinkRequest, error) {
	l := model.RegisterLinkRequest{}
	json.NewDecoder(r.Body).Decode(&l)
	if l.ThreshHold == nil || l.Url == nil || l.Method == nil || !ValidateMethod(*l.Method) {
		return l, Err.ErrMissingFields
	}
	return l, nil
}

func ParseLinkGet(r *http.Request) (int64, error) {
	l := model.LinkRequest{}
	json.NewDecoder(r.Body).Decode(&l)
	if l.LinkId == nil {
		return 0, Err.ErrMissingFields
	}
	return *l.LinkId, nil
}
