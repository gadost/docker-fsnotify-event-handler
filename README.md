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
docker run -it --privileged \n
 --network=host \n
 -v /var/run/docker.sock:/var/run/docker.sock \n
 -v /root/gotest:/dir \n
 -e WATCH_PATH=/dir \n
 -e "REDIS_ADDR=somehost:6379" \n
 maxn/docker-fsnotify-event-handler:v1
```
