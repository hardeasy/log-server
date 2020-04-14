package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/olivere/elastic"
	"log"
	"log-server/config"
	"time"
)


type Dao struct {
	Cfg *config.Config
	Db *gorm.DB
	Redis *redis.Client
	Elastic *elastic.Client
	LogDao *LogDao
	UserDao *UserDao
	PushappDao *PushappDao
}

func New(cfg *config.Config) *Dao {
	d := &Dao{
		Cfg:cfg,
		Db: newDb(cfg.Mysql),
		Redis: newRedis(cfg.Redis),
		Elastic: newElasticSearch(cfg.ElasticSearch),
		LogDao: &LogDao{},
		UserDao: &UserDao{},
		PushappDao: &PushappDao{},
	}
	//d.db.LogMode(true)
	return d
}

func newDb(cfg *config.Mysql) *gorm.DB{
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Passwd, cfg.Host, cfg.Port, cfg.DbName))
	if err != nil {
		panic(err)
	}
	return db
}

func newRedis(cfg *config.Redis) *redis.Client{
	rds := redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Passwd,
		DB:                 cfg.Db,
	})
	return rds
}

func newElasticSearch(cfg *config.ElasticSearch) *elastic.Client {
	cxt,cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL([]string{cfg.Url}...))
	if err != nil {
		panic(err)
	}
	defer cancel()
	info, code, err := client.Ping(cfg.Url).Do(cxt)
	if err != nil {
		panic(err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	esversion, err := client.ElasticsearchVersion(cfg.Url)
	if err != nil {
		// Handle error
		panic(err)
	}
	log.Printf("Elasticsearch version %s\n", esversion)
	return client
}