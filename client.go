package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func client(command, key, value *string, n, parallel int) {
	conn, err := makeConn()
	if err != nil {
		log.Panicln("[ERROR] cannot make connection; error:", err)
	}
	defer conn.Close()

	switch *command {
	case Get:
		send(get(key), conn)
	case Set:
		send(set(key, value), conn)
	case Remove:
		send(remove(key), conn)
	case Benchmark:
		benchmark(n, parallel)
	case Repl:
		repl(conn)
	}
}

func makeConn() (*Conn, error) {
	serverAddress, err := net.ResolveTCPAddr("tcp", "localhost:2934")
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, serverAddress)
	if err != nil {
		return nil, err
	}
	return NewConn(conn), nil
}

func repl(conn *Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			continue
		} else if isPrefix {
			os.Exit(-1)
		}
		execute(string(line), conn)
	}
}

func execute(line string, conn *Conn) {
	tokens := strings.Split(line, " ")
	if len(tokens) < 1 {
		log.Println("empty command")
		return
	}
	switch tokens[0] {
	case Get:
		if len(tokens) != 2 {
			log.Println("illegal get")
			return
		}
		send(get(&tokens[1]), conn)
	case Set:
		if len(tokens) != 3 {
			log.Println("illegal set")
			return
		}
		send(set(&tokens[1], &tokens[2]), conn)
	case Remove:
		if len(tokens) != 2 {
			log.Println("illegal remove")
			return
		}
		send(remove(&tokens[1]), conn)
	case Benchmark:
		if len(tokens) != 3 {
			log.Println("illegal benchmark")
			return
		}
		n, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Printf("%s is not a number", tokens[1])
			return
		}
		parallel, err := strconv.Atoi(tokens[2])
		if err != nil {
			log.Printf("%s in not a number", tokens[2])
			return
		}
		benchmark(n, parallel)
	default:
		log.Println("don't know how to handle", tokens[0])
	}
}

type bench struct {
	sum time.Duration
	avg time.Duration
	p50 time.Duration
	p90 time.Duration
	p99 time.Duration
}

func benchmark(n, parallel int) {
	randomStrings := makeRandomPairPool()
	conns := make([]*Conn, parallel)
	for i := 0; i < parallel; i++ {
		conn, err := makeConn()
		if err != nil {
			log.Println("[ERROR] cannot establish connection; error:", err)
			return
		}
		conns[i] = conn
	}

	var wg sync.WaitGroup
	elapsedTimes := make([][]time.Duration, parallel)
	for i := 0; i < parallel; i += 1 {
		elapsedTimes[i] = make([]time.Duration, n)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < n; j += 1 {
				start := time.Now()
				result := send(set(&randomStrings[rand.Intn(256)], &randomStrings[rand.Intn(256)]), conns[i])
				elapsed := time.Since(start)
				if result.Type == Failure {
					fmt.Printf("[ERROR] worker %d, error %s\n", i, result)
				}
				elapsedTimes[i][j] = elapsed
			}
		}(i)
	}
	wg.Wait()

	result := make([]bench, parallel)
	for i := 0; i < parallel; i += 1 {
		sum := time.Duration(0)
		sort.Slice(elapsedTimes[i], func(k, l int) bool {
			return elapsedTimes[i][k] < elapsedTimes[i][l]
		})
		for j := 0; j < n; j += 1 {
			sum += elapsedTimes[i][j]
		}
		result[i].sum = sum
		result[i].avg = time.Duration(float64(sum) / float64(n))
		result[i].p50 = elapsedTimes[i][n/2]
		result[i].p90 = elapsedTimes[i][int(math.Round(float64(n)*0.9))]
		result[i].p99 = elapsedTimes[i][int(math.Round(float64(n)*0.99))]
	}

	fmt.Println("-> Benchmark result")
	for i := 0; i < parallel; i += 1 {
		fmt.Println("# Worker", i)
		fmt.Println("    sum:", result[i].sum)
		fmt.Println("    avg:", result[i].avg)
		fmt.Println("    p50:", result[i].p50)
		fmt.Println("    p90:", result[i].p90)
		fmt.Println("    p99:", result[i].p99)
	}
}

func makeRandomPairPool() (strs []string) {
	rand.Seed(time.Now().UnixNano())
	strs = make([]string, 256)
	for i := 0; i < 256; i += 1 {
		strs[i] = makeRandomString()
	}
	return
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetLen = len(charset)

func makeRandomString() string {
	buffer := make([]byte, 128)
	for i := range buffer {
		buffer[i] = charset[rand.Intn(charsetLen)]
	}
	return string(buffer)
}

func get(key *string) *Message {
	return &Message{Type: Get, Key: *key}
}

func set(key *string, value *string) *Message {
	return &Message{Type: Set, Key: *key, Value: *value}
}

func remove(key *string) *Message {
	return &Message{Type: Remove, Key: *key}
}

func send(message *Message, conn *Conn) *Message {
	err := conn.Send(message)
	if err != nil {
		log.Println("[ERROR] error occurred while sending; error:", err)
		return nil
	}
	result, err := conn.Receive()
	if err != nil {
		log.Println("[ERROR] error occurred while receiving; error:", err)
		return nil
	}

	return result
}
