package main

import (
	"flag"
	"log"
)

func main() {
	mode := flag.String("mode", "", "server or client")
	key := flag.String("key", "", "key string")
	value := flag.String("value", "", "value string")
	command := flag.String("command", "", "[get|set|remove]")
	flag.Parse()

	switch *mode {
	case "server":
		server()
	case "client":
		client(command, key, value)
	default:
		log.Panicln("[ERROR] don't know how to run", mode)
	}
}
