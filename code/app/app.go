package app

import (
	"monitor/config"
	"monitor/router"
	"net/http"
)

func Run() {
	mux := router.Router()
	server := &http.Server{
		Addr:         config.Addr,
		IdleTimeout:  config.IdleTimeOut,
		ReadTimeout:  config.ReadTimeOut,
		WriteTimeout: config.WriteTimeOut,
		Handler:      mux,
	}
	go monitor()
	server.ListenAndServe()
}
