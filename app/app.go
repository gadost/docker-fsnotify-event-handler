package app

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis"
)

var c Config

func Start() {
	c.MarshalConfig()
	//define watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	client := c.Client()
	c.RegisterAgent(client)
	go c.Walking(watcher)
	go QueueProcessor(client)
	c.CheckEvent(watcher, client)
}

func (c *Config) Walking(watcher *fsnotify.Watcher) {
	_ = filepath.Walk(c.Path, func(path string, f os.FileInfo, err error) error {
		watcher.Add(path)
		return nil
	})
}

func (c Config) CheckEvent(watcher *fsnotify.Watcher, client *redis.Client) {
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fi, err := os.Stat(event.Name)
				if err != nil {
					break
				}
				if fi.Mode().IsDir() {
					watcher.Add(event.Name)
					break
				}
				if AllowedEvents(event.Op.String()) {
					log.Println("event:", event)
					AddToQueue(c.AgentName, "wait", client)
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
		go CheckQueue(c.AgentName, "wait", client)
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
