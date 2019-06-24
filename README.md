## overview
docker-fsnotify-event-handler developed for processing some filesystem events like letsencrypt cert renew for one domain on multiply hosts ( f.e. on my case - N nginx instances on different hosts using same certs. On certs renew le companion reloading only one nginx, other all proceed using old certs.).

## Install

```bash
go get github.com/gadost/docker-fsnotify-event-handler
cd $GOPATH/src/github.com/gadost/docker-fsnotify-event-handler
go build .
export WATCH_PATH=/path/to/watch ; ./docker-fsnotify-event-handler

```
## Struct
```golang
RedisAddr  string  `env:"REDIS_ADDR" envDefault:"localhost:6379"`
RedisPass  string  `env:"REDIS_PASS" envDefault:""`
RedisDB    int     `env:"REDIS_DB" envDefault:"0"`
AgentsSet  string  `env:"AGENTS_SET_NAME" envDefault:"agents"`
QueueName  string  `env:"QUEUE_NAME" envDefault:"lecc"`
AgentName  string  `env:"HOSTNAME"`
```

## Docker example

```bash
docker run -it --privileged --network=host -v /var/run/docker.sock:/var/run/docker.sock -v /root/gotest:/dir -e WATCH_PATH=/dir -e "REDIS_ADDR=somehost:6379" maxn/docker-fsnotify-event-handler:v1
```

## build
```bash
docker build .
```
