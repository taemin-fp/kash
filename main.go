package main

import (
	"flag"
)

func main() {
	mode := flag.String("mode", "", "server or client")
	initialCapacity := flag.Int("capacity", 4096000, "(server) initial memory allocation capacity")
	workers := flag.Int("workers", 32, "(server) worker pool size")
	key := flag.String("key", "", "(client) key string")
	value := flag.String("value", "", "(client) value string")
	command := flag.String("command", "", "(client) [get|set|remove|repl]")
	n := flag.Int("n", 0, "(benchmark) benchmark running count")
	parallel := flag.Int("parallel", 0, "(benchmark) benchmark parallelism")
	keySize := flag.Int("key-size", 128, "(benchmark) benchmark key size")
	flag.Parse()

	switch *mode {
	case "server":
		server(*initialCapacity, *workers)
	case "client":
		client(command, key, value, *n, *parallel, *keySize)
	default:
		flag.Usage()
	}
}
