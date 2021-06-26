package main

import (
	"io"
	"log"
	"net"
	"sync"
)

type Handler struct {
	storage Storage
	mutex   sync.RWMutex
}

type messageAndConn struct {
	message *Message
	conn    *Conn
}

func server() {
	handler := Handler{storage: GetStorage(), mutex: sync.RWMutex{}}

	l, err := net.Listen("tcp", ":2934")
	if err != nil {
		log.Panicln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("[WARN] cannot accept; error:", err)
			continue
		}

		clientConn := NewConn(conn)
		go readMessage(clientConn, &handler)
	}
}

func readMessage(conn *Conn, handler *Handler) {
	for {
		message, err := conn.Receive()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println("[WARN] error on readMessage; error:", err)
			return
		}

		go handler.handle(&messageAndConn{message: message, conn: conn})
	}
}

func (h *Handler) respond(conn *messageAndConn) {
	err := conn.conn.Send(conn.message)
	if err != nil {
		log.Println("[WARN] error on responding; error: ", err)
	}
}

func (h *Handler) handle(conn *messageAndConn) {
	message := conn.message

	var result *Message
	switch message.Type {
	case Get:
		h.mutex.RLock()
		defer h.mutex.RUnlock()
		value := h.storage.Get(message.Key)
		result = &Message{Type: Success, Value: value}
	case Set:
		h.mutex.Lock()
		defer h.mutex.Unlock()
		if message.Value == nil || message.Value == "" {
			result = &Message{Type: Failure}
		} else {
			h.storage.Set(message.Key, message.Value)
			result = &Message{Type: Success}
		}
	case Remove:
		h.mutex.Lock()
		defer h.mutex.Unlock()
		if message.Key == "" || h.storage.Get(message.Key) == nil {
			result = &Message{Type: Failure}
		} else {
			h.storage.Remove(message.Key)
			result = &Message{Type: Success}
		}
	default:
		log.Println("[INFO] unsupported message", message)
	}
	go h.respond(&messageAndConn{message: result, conn: conn.conn})
}
