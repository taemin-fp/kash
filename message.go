package main

const (
	Get       = "get"
	Set       = "set"
	Remove    = "remove"
	Benchmark = "benchmark"
	Repl      = "repl"
	Success   = "success"
	Failure   = "failure"
)

type Message struct {
	Type  string
	Key   string
	Value interface{}
}
