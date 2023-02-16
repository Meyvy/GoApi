package view

import (
	"encoding/json"
	"monitor/db"
	Err "monitor/error"
	"monitor/model"
	"monitor/util"
	"net/http"
	"strconv"
)

func RegisterLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrHandle(Err.ErrNotFound).ServeHTTP(w, r)
		return
	}
	l, err := util.ParseRegisterLink(r)
	if err != nil {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	user_id, _ := strconv.ParseInt(r.Header.Get("User_id"), 10, 64)
	id, err := db.AddLink(l, user_id)
	if err != nil {
		ErrHandle(Err.ErrInternal).ServeHTTP(w, r)
		return
	}
	json.NewEncoder(w).Encode(model.RegisterLinkResponse{
		Status: "success",
		LinkId: id,
	})
}

func FetchAllLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrHandle(Err.ErrNotFound).ServeHTTP(w, r)
		return
	}
	userId, _ := strconv.ParseInt(r.Header["User_id"][0], 10, 64)
	list, err := db.GetAllLink(userId)
	if err != nil {
		ErrHandle(Err.ErrInternal).ServeHTTP(w, r)
		return
	}
	json.NewEncoder(w).Encode(list)
}

func FetchLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrHandle(Err.ErrNotFound).ServeHTTP(w, r)
		return
	}
	linkId, err := util.ParseLinkGet(r)
	if err != nil {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	userId, _ := strconv.ParseInt(r.Header["User_id"][0], 10, 64)
	link, err := db.GetLink(userId, linkId)

	if err != nil {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	requests, err := db.GetTodayRequest(userId, linkId)
	if err != nil {
		ErrHandle(Err.ErrInternal).ServeHTTP(w, r)
		return
	}
	res := model.LinkResponse{
		Link:     link,
		Requests: requests,
	}
	json.NewEncoder(w).Encode(res)
}
