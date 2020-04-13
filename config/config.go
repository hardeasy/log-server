package config

import (
	"flag"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
	Conf *Config
)

const (
	Application_DEV = "dev"
	Application_TEST = "test"
	Application_PROD = "prod"
)

func init() {
	flag.StringVar(&configPath, "conf", "conf.example.toml", "default conf path")
}

func Init() error {
	Conf = Default()
	_, err := toml.DecodeFile(configPath, &Conf)
	return err
}

func Default() *Config{
	return &Config{
		Application: &Application{
			Env:     "dev",
			LogPath: "log.log",
			PidPath: "pid.pid",
		},
		HTTPServer: &HTTPServer{Addr:":8080"},
		Mysql:      &Mysql{},
		Redis: &Redis{},
		PushServer: &PushServer{},
		ElasticSearch: &ElasticSearch{},
	}
}

type Config struct {
	Application *Application
	HTTPServer *HTTPServer
	Mysql *Mysql
	Redis *Redis
	PushServer *PushServer
	ElasticSearch *ElasticSearch
}

type Application struct {
	Env string
	LogPath string
	PidPath string
}

type HTTPServer struct {
	Addr string
}

type Mysql struct {
	Host string
	Port int
	User string
	Passwd string
	DbName string
}

type Redis struct {
	Host string
	Port int
	Passwd string
	Db int
}

type PushServer struct {
	Token string
	QueueKey string
}

type ElasticSearch struct {
	Url string
}