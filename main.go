package main


import (
    "github.com/fsnotify/fsnotify"
    "log"
    "os"
    "path/filepath"
    "github.com/gadost/docker-fsnotify-event-handler/sqs"
    "github.com/go-redis/redis"
    "time"
)

var c = sqs.MarshalConfig()

func main() {
    //define watcher
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    client := sqs.Client(c)
    sqs.RegisterAgent(c.AgentsSet, c.AgentName, client)
    go Walking(watcher)
    go QueueProcessor(client)
    CheckEvent(watcher, client,c)
}

func Walking(watcher *fsnotify.Watcher) {
    _ = filepath.Walk(os.Getenv("WATCH_PATH"), func(path string, f os.FileInfo, err error) error {
        watcher.Add(path)
        return nil
    })
}

func CheckEvent(watcher *fsnotify.Watcher, client *redis.Client,c sqs.Config) {
    done := make(chan bool)
    go func() {
        for {
            select {
            case event := <-watcher.Events:
                fi,err := os.Stat(event.Name)
                if err != nil {
                    break
                }
                if fi.Mode().IsDir() {
                    watcher.Add(event.Name)
                    break
                }
                if  AllowedEvents(event.Op.String()) {
                    log.Println("event:", event)
                    sqs.AddToQueue(c.AgentName, "wait", client)
                }
            case err := <-watcher.Errors:
                log.Println("error:", err)
            }
        }
    }()
    <-done
}

func QueueProcessor(client *redis.Client) {
  for {
    <-time.After(time.Duration(c.Interval) * time.Second)
    go sqs.CheckQueue(c.AgentName, "wait", client)
  }
}

func AllowedEvents(e string) bool {
    switch e {
    case
    "RENAME",
    "CHMOD",
    "CREATE",
    "WRITE":
    return true
    }
    return false
}
