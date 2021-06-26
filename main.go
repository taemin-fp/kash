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
	n := flag.Int("n", 0, "benchmark running count")
	parallel := flag.Int("parallel", 0, "benchmark parallelism")
	keySize := flag.Int("key-size", 128, "benchmark key size")
	flag.Parse()

	switch *mode {
	case "server":
		server()
	case "client":
		client(command, key, value, *n, *parallel, *keySize)
	default:
		log.Panicln("[ERROR] don't know how to run", mode)
	}
}
