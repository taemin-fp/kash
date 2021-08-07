package main

import (
	"flag"
)

type arguments struct {
	mode            string
	initialCapacity int
	workers         int
	key             string
	value           string
	command         string
	n               int
	parallel        int
	keySize         int
}

func main() {
	args := getArguments()

	switch args.mode {
	case "server":
		server(args.initialCapacity, args.workers)
	case "client":
		client(&args.command, &args.key, &args.value, args.n, args.parallel, args.keySize)
	default:
		flag.Usage()
	}
}

func getArguments() arguments {
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

	return arguments{
		*mode,
		*initialCapacity,
		*workers,
		*key,
		*value,
		*command,
		*n,
		*parallel,
		*keySize,
	}
}
