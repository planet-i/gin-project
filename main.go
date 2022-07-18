package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	_ "github.com/go-sql-driver/mysql"
	"github.com/planet-i/gin-project/models"
	"github.com/planet-i/gin-project/pkg/logging"
	"github.com/planet-i/gin-project/pkg/setting"
	"github.com/planet-i/gin-project/routers"
)

func main() {
	// 为了控制各个包配置项的初始化顺序
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
