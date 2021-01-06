package main

import (
	"fmt"
	"log"
	"net/http"

	"gin_tutorial/pkg/setting"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	config, err := setting.NewConfig()
	if err != nil {
		log.Fatal("Can init Config with Enviroment:", err)
	}
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", config.HTTPPort),
		Handler: router,
		ReadTimeout: config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}