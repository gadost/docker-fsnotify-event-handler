package main


import (
    "github.com/fsnotify/fsnotify"
    "log"
    "os"
    "path/filepath"
)

func main() {
    //define watcher
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()
    //Add watcher start dir
    go Walking(watcher)
    CheckEvent(watcher)
}

func Walking(watcher *fsnotify.Watcher) {
    _ = filepath.Walk(os.Getenv("WATCH_PATH"), func(path string, f os.FileInfo, err error) error {
        watcher.Add(path)
        return nil
    })
}

func CheckEvent(watcher *fsnotify.Watcher) {
    done := make(chan bool)
    go func() {
        for {
            select {
            case event := <-watcher.Events:
                fi,_ := os.Stat(event.Name)
                if fi.Mode().IsDir() {
                    watcher.Add(event.Name)
                    break
                }
                log.Println("event:", event)
                if  string(event.Op.String()) == "RENAME" || string(event.Op.String()) == "CHMOD" || string(event.Op.String()) == "CREATE" {
                    e, err := os.Create(os.Getenv("EVENT_MARKER"))
                    if err != nil {
                        log.Fatal(err)
                    }
                    log.Println(e)
                    e.Close()
                }
            case err := <-watcher.Errors:
                log.Println("error:", err)
            }
        }
    }()
    <-done
}
