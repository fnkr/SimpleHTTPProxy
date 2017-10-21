package main

import (
    "flag"
)

type Config struct {
    Host string
    Port uint16
    Verbose bool
}

func getConfig() Config {
    host := flag.String("host", "", "")
    port := flag.Int("port", int(1080), "")
    verbose := flag.Bool("verbose", false, "")

    flag.Parse()

    return Config{
        Host: *host,
        Port: uint16(*port),
        Verbose: *verbose,
    }
}
