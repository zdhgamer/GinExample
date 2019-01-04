package setting

import (
	"github.com/go-ini/ini"
	"time"
	"log"
)

var(
	Cfg *ini.File
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	PageSize int
	JwtSecret string
)

func init()  {
	var err error
	Cfg,err = ini.Load("src/conf/app.ini")
	if err!=nil{
		log.Fatal("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase()  {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer()  {
	sec,err := Cfg.GetSection("server")
	if err!=nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp()  {
	sec,err := Cfg.GetSection("app")
	if err!=nil {
		log.Fatal("Fail to get section 'app': %v", err)
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = sec.Key("JWT_SECRET").MustString("23347$040412")
}