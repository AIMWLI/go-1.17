package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize  string
	JwtSecret string
}

var AppSetting = &App{}

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

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Probes struct {
	KV map[string]string
}

var ProbesSetting = &Probes{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting#setup fail to parse conf/app.ini: %v", err)
	}
	loadApp()
	loadServer()
	loadDataBase()
	loadRedis()
	loadProbes()
}

func loadProbes() {
	section := cfg.Section("probes")
	ProbesSetting.KV = section.KeysHash()
}

func loadRedis() {
	mapTo("redis", RedisSetting)
}

func loadApp() {
	mapTo("app", AppSetting)
}
func loadServer() {
	mapTo("server", ServerSetting)
}
func loadDataBase() {
	mapTo("database", DataBaseSetting)
	log.Printf("", DataBaseSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("setting#mapTo section [%v] err: %v", section, err)
	}

}
