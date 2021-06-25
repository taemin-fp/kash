package main

import (
	"bytes"
	"encoding/gob"
	"io"
	"log"
	"net"
	"sync"
)

type Handler struct {
	storage Storage
	mutex sync.Mutex
}

func server() {
	handler := Handler{ storage: GetStorage() }

	l, err := net.Listen("tcp", ":2934")
	if err != nil {
		log.Panicln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("[WARN] cannot accept", conn)
			continue
		}

		go readMessage(conn, &handler)
	}
}

func readMessage(conn net.Conn, handler *Handler) {
	var codecBuffer bytes.Buffer
	message := Message{}
	decoder := gob.NewDecoder(&codecBuffer)
	recvbuf := make([]byte, 8192)

	for {
		n, err := conn.Read(recvbuf)
		if err != nil {
			if err == io.EOF {
				log.Println("[INFO] connection closed by peer", conn.RemoteAddr().String())
				return
			}
			log.Println("[WARN] cannot receive data; error:", err)
			return
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

	go handler.handle(&message, conn)
}

func (h *Handler) handle(message *Message, conn net.Conn) {
	var result *Message

	h.mutex.Lock()
	switch message.Type {
	case Set:
		log.Printf("[INFO] key %s, value %s", message.Key, message.Value)
		if message.Value == nil || message.Value == "" {
			result = &Message{Type: Failure}
		} else {
			h.storage.Set(message.Key, message.Value)
			result = &Message{Type: Success}
		}
	case Get:
		value := h.storage.Get(message.Key)
		result = &Message{Type: Success, Value: value}
	default:
		log.Println("[INFO] unsupported message {}", message)
	}
	h.mutex.Unlock()

	go h.response(result, conn)
}

func (h *Handler) response(result *Message, conn net.Conn) {
	buffer := serialize(result)

	n, err := conn.Write(buffer)
	if n != len(buffer) {
		log.Printf("[WARN] cannot send response completely; %d bytes sent", n)
		return
	}
	if err != nil {
		log.Println("[WARN] cannot send response")
	}
}
