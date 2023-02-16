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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrHandle(Err.ErrNotFound).ServeHTTP(w, r)
		return
	}
	s, err := util.ParseUserSignUp(r)
	if err != nil {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	ok := util.ValidateUser(s)
	if !ok {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	id, err := db.AddUser(s)
	if err != nil {
		ErrHandle(Err.ErrInternal).ServeHTTP(w, r)
		return
	}
	json.NewEncoder(w).Encode(model.RegisterUserResponse{
		Status: "success",
		UserId: id,
	})
}

func UserWarningAll(w http.ResponseWriter, r *http.Request) {
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
	new_list := make([]model.Link, 0)
	for _, v := range list {
		if v.Failures >= v.ThreshHold {
			new_list = append(new_list, v)
		}
	}
	json.NewEncoder(w).Encode(new_list)
}
