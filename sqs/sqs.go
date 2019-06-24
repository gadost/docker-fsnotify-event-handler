package sqs

import (
    "github.com/go-redis/redis"
    "log"
    "github.com/caarlos0/env"
    "github.com/gadost/docker-fsnotify-event-handler/event"
)


type Config struct {
	redisAddr  string  `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	redisPass  string  `env:"REDIS_PASS" envDefault:""`
	redisDB    int     `env:"REDIS_DB" envDefault:"0"`
    AgentsSet  string  `env:"AGENTS_SET_NAME" envDefault:"agents"`
    QueueName  string  `env:"QUEUE_NAME" envDefault:"lecc"`
    AgentName  string  `env:"HOSTNAME"`
}

var client *redis.Client

func init(){}
// Define Client Redis from config
func Client(c Config) *redis.Client {
    client = redis.NewClient(&redis.Options{
        Addr:     c.redisAddr,
        Password: c.redisPass,
        DB:       c.redisDB,
    })
    _, err := client.Ping().Result()
    if err != nil {
        log.Println("Redis server unreachable! Exit..")
        panic(err)
    }
    return client
}

func MarshalConfig() Config {
    cfg := Config{}
    if err := env.Parse(&cfg); err != nil {
        log.Printf("%+v\n", err)
    }
    return cfg
}

func RegisterAgent(k string,v string, client *redis.Client) {
    err := client.SAdd(k,v)
    log.Println(err)
}

func QueueComplited(k string, v string, client *redis.Client) {
    err := client.Set(k, v, 0).Err()
    if err != nil {
        log.Print(err)
    }
}
func AddToQueue(k string, v string, client *redis.Client) {
    err := client.Set(k, v, 0).Err()
    if err != nil {
        log.Print(err)
    }
}

func CheckQueue(k string,v string, client *redis.Client) {
    val, err := client.Get(k).Result()
    if err != nil {
        log.Print(err)
    } else {
        if val == v {
            event.Handler()
            QueueComplited(k, "done", client)
        }
    }
}
