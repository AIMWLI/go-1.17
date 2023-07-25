package setting

import (
	"github.com/go-ini/ini"

	"log"
	"time"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Timeout      time.Duration
}

var ServerSetting = &Server{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting#setup fail to parse conf/app.ini: %v", err)
	}
	mapTo("server", ServerSetting)
	//mapTo("app", )
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("setting#mapTo section [%v] err: %v", section, err)
	}

}
