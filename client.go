package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
)

func client(command, key, value *string) {
	switch *command {
	case Get:
		send(get(key))
	case Set:
		send(set(key, value))
	case Remove:
		send(remove(key))
	}
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

func send(message *Message) {
	serverAddress, err := net.ResolveTCPAddr("tcp", "localhost:2934")
	if err != nil {
		log.Panicln(err)
	}

	conn, err := net.DialTCP("tcp", nil, serverAddress)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	buffer := serialize(message)
	n, err := conn.Write(buffer)
	if n != len(buffer) {
		log.Printf("[WARN] cannot send message completely; sent %d bytes", n)
		return
	} else if err != nil {
		log.Panicln(err)
	}

	result := receive(conn)
	switch result.Type {
	case Success:
		switch message.Type {
		case Get:
			fmt.Println(result.Value)
		default:
			log.Println("[WARN] unknown result", result.Type)
		}
	case Failure:
		log.Println("[ERROR] failure")
	default:
		log.Println("[WARN] unknown message", result)
	}
}

func receive(conn net.Conn) *Message {
	var codecBuffer bytes.Buffer
	message := Message{}
	decoder := gob.NewDecoder(&codecBuffer)
	recvbuf := make([]byte, 8192)

	for {
		n, err := conn.Read(recvbuf)
		if err != nil {
			if err == io.EOF {
				log.Println("[INFO] connection closed by peer", conn.RemoteAddr().String())
				return nil
			}
			log.Println("[WARN] cannot receive data; error:", err)
			return nil
		}

		if n > 0 {
			data := recvbuf[:n]
			codecBuffer.Write(data)

			if err := decoder.Decode(&message); err != nil {
				log.Println("[WARN] cannot decode message; error:", err)
				continue
			}
			break
		}
	}

	return &message
}
