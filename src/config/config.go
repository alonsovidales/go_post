package config

import (
    "code.google.com/p/gcfg"
    "log"
)

type Config struct {
    Default struct {
        Port int
        Ip string
        Logging string
    }
    Auth struct {
        Active bool
        Url string
    }
    Redis struct {
        Server []string
    }
}

var cfg *Config = nil

func Get() *Config {
    if cfg != nil {
        return cfg
    }

    var config Config
    err := gcfg.ReadFileInto(&config, "etc/go_post.gcfg")

    if err != nil {
        err = gcfg.ReadFileInto(&config, "/etc/go_post.gcfg")
        if err != nil {
            log.Fatal("Can't read the configuration from etc/go_post.gcfg or /etc/go_post.gcfg: ", err)
        }
    }
    cfg = &config

    return cfg
}
