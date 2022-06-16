# docker-fsnotify-event-handler

## Overview

docker-fsnotify-event-handler developed for processing some filesystem events like letsencrypt cert renew for one domain on multiply hosts ( f.e. on my case - N nginx instances on different hosts using same certs. On certs renew le companion reloading only one nginx, other all proceed using old certs.).

## Install

```bash
git clone github.com/gadost/docker-fsnotify-event-handler
cd docker-fsnotify-event-handler
go build .
export WATCH_PATH=/path/to/watch ; ./docker-fsnotify-event-handler

```

## Environment

```golang
// redis host , pass , db
RedisAddr string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
RedisPass string `env:"REDIS_PASS" envDefault:""`
RedisDB   int    `env:"REDIS_DB" envDefault:"0"`

// agent set 
AgentsSet string `env:"AGENTS_SET_NAME" envDefault:"agents"`
QueueName string `env:"QUEUE_NAME" envDefault:"lecc"`
AgentName string `env:"HOSTNAME"`

//check interval
Interval  int    `env:"INTERVAL" envDefault:"20"`

// Path to watch
Path      string `env:"WATCH_PATH"`

// Docker image "test/image"
Image     string `env:"IMAGE"`

// execute command when event happens ( f.e. "nginx -s reload" )
Command   string `env:"COMMAND"`

```

## Docker example

```bash
docker run -it --privileged --network=host -v /var/run/docker.sock:/var/run/docker.sock -v /path/to/watching/dir:/dir -e WATCH_PATH=/dir -e "REDIS_ADDR=somehost:6379" maxn/docker-fsnotify-event-handler:latest
```

## Build

```bash
docker build .
```
