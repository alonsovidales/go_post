package main

import (
    "handlers"
)

func main() {
    dispatcher := handlers.Dispatcher()
    dispatcher.Run()

    /*c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10) * time.Second)

    defer c.Close()

    s, err := c.Cmd("info").Str()
    fmt.Println("mykey0:", s, err)

    log.Fatal("MUAHAHAHAH")*/
}
