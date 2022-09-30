package setting

import (
	"log"
	"time"

	"github.com/xppcnn/gin-demo/config"
)

var (
	Cfg          config.Config
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg = config.GetConfig()
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.RunMode
}

func LoadServer() {
	sec := Cfg.Server
	HTTPPort = sec.HttpPort
	ReadTimeout = time.Duration(sec.ReadTimeout) * time.Second
	WriteTimeout = time.Duration(sec.WriteTimeout) * time.Second
}

func LoadApp() {
	sec := Cfg.App
	JwtSecret = sec.JwtSecret
	PageSize = sec.PageSize
}
