package handlers

import (
    "config"
    "fmt"
    "logger"
    "net/http"
)

var port = config.Get().Default.Port

type dispatcher struct {
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, "Hello world, from Go! ", r.Method, r.FormValue("pepe"))
}

func Dispatcher() *dispatcher {
    dispatcher := dispatcher{};

    return &dispatcher
}

func (self *dispatcher) Run() bool {
    http.HandleFunc("/hello", hello)

    logger.Info("Server listening on port: %d", port)
    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    if err != nil {
        fmt.Print(fmt.Sprintf(":%d\n", port))
        logger.Fatal("Can't listen on port: %d", port)
    }

    return true
}
