package main

import (
	"fmt"
	"net/http"

	"github.com/planet-i/gin-project/pkg/setting"
	"github.com/planet-i/gin-project/routers"
)

func main() {
	// 1.创建路由
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
