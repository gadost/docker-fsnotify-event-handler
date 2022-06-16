package app

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/go-redis/redis"
)

type Config struct {
	RedisAddr string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	RedisPass string `env:"REDIS_PASS" envDefault:""`
	RedisDB   int    `env:"REDIS_DB" envDefault:"0"`
	AgentsSet string `env:"AGENTS_SET_NAME" envDefault:"agents"`
	QueueName string `env:"QUEUE_NAME" envDefault:"lecc"`
	AgentName string `env:"HOSTNAME"`
	Interval  int    `env:"INTERVAL" envDefault:"20"`
	Path      string `env:"WATCH_PATH"`
	Command   string `env:"COMMAND"`
	Image     string `env:"IMAGE"`
}

var rClient *redis.Client

// Define Client Redis from config
func (c *Config) Client() *redis.Client {
	rClient = redis.NewClient(&redis.Options{
		Addr:     c.RedisAddr,
		Password: c.RedisPass,
		DB:       c.RedisDB,
	})
	_, err := rClient.Ping().Result()
	if err != nil {
		log.Println("Redis server unreachable! Exit..")
		panic(err)
	}
	return rClient
}

func (c *Config) MarshalConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return &cfg
}

func (c *Config) RegisterAgent(rClient *redis.Client) {
	err := rClient.SAdd(c.AgentsSet, c.AgentName)
	log.Println(err)
}

func QueueComplited(k string, v string, rClient *redis.Client) {
	err := rClient.Set(k, v, 0).Err()
	if err != nil {
		log.Print(err)
	}
}
func AddToQueue(k string, v string, rClient *redis.Client) {
	err := rClient.Set(k, v, 0).Err()
	if err != nil {
		log.Print(err)
	}
}

func CheckQueue(k string, v string, rClient *redis.Client) {
	val, err := rClient.Get(k).Result()
	if err != nil {
		log.Print(err)
	} else {
		if val == v {
			Handler()
			QueueComplited(k, "done", rClient)
		}
	}
}
