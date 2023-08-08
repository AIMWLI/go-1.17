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

type Database struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

var DataBaseSetting = &Database{}

var cfg *ini.File

func Setup() {
	loadServer()
	//mapTo("app", )
}
func loadServer() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting#setup fail to parse conf/app.ini: %v", err)
	}
	mapTo("server", ServerSetting)
}
func load() {

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("setting#mapTo section [%v] err: %v", section, err)
	}

}
