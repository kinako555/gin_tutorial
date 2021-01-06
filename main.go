package main

import (
	"fmt"
	"log"
	"net/http"

	"gin_tutorial/pkg/setting"
	"gin_tutorial/routers"
)

func main() {
	config, err := setting.NewConfig()
	if err != nil {
		log.Fatal("Can init Config with Enviroment:", err)
	}
	router := routers.InitRouter(config.AppConfig)
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", config.HTTPPort),
		Handler: router,
		ReadTimeout: config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}