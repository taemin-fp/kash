package main

import (
	"flag"
	"log"
)

func main() {
	mode := flag.String("mode", "", "server or client")
	initialCapacity := flag.Int("capacity", 4096000, "initial memory allocation capacity")
	workers := flag.Int("workers", 32, "worker pool size")
	key := flag.String("key", "", "key string")
	value := flag.String("value", "", "value string")
	command := flag.String("command", "", "[get|set|remove]")
	n := flag.Int("n", 0, "benchmark running count")
	parallel := flag.Int("parallel", 0, "benchmark parallelism")
	keySize := flag.Int("key-size", 128, "benchmark key size")
	flag.Parse()

	switch *mode {
	case "server":
		server(*initialCapacity, *workers)
	case "client":
		client(command, key, value, *n, *parallel, *keySize)
	default:
		log.Panicln("[ERROR] don't know how to run", mode)
	}
}
