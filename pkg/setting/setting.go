// 此模块用来加载配置
package setting

import (
	"time"

	"github.com/go-ini/ini"
	"github.com/planet-i/gin-project/pkg/logging"
)

// 1.定义对应变量
var (
	Cfg *ini.File

	RunMode string

	PageSize  int
	JwtSecret string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	var err error
	// 2.加载配置文件
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		logging.Fatal("Fail to parse 'conf/app.ini': %v", err)
	}
	// 3. 将配置文件中的对应数据赋值到变量中
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		logging.Fatal("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		logging.Fatal("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
