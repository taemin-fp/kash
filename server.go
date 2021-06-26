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

func server(initialCapacity int, workerPoolSize int) {
	handler := Handler{storage: GetStorage(initialCapacity), mutex: sync.RWMutex{}}
	in := make(chan *messageAndConn)
	out := make(chan *messageAndConn)

	for i := 0; i < workerPoolSize; i += 1 {
		go handler.handler(in, out)
		go handler.responder(out)
	}

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
		go readMessage(clientConn, in)
	}
}

func readMessage(conn *Conn, in chan<- *messageAndConn) {
	for {
		message, err := conn.Receive()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println("[WARN] error on readMessage; error:", err)
			return
		}

		in <- &messageAndConn{message: message, conn: conn}
	}
}

func (h *Handler) responder(out <-chan *messageAndConn) {
	for res := range out {
		err := res.conn.Send(res.message)
		if err != nil {
			log.Println("[WARN] error on responding; error: ", err)
		}
	}
}

func (h *Handler) handler(in <-chan *messageAndConn, out chan<- *messageAndConn) {
	for req := range in {
		message := req.message

		var result *Message
		switch message.Type {
		case Get:
			h.mutex.RLock()
			value := h.storage.Get(message.Key)
			h.mutex.RUnlock()
			result = &Message{Type: Success, Value: value}
		case Set:
			if message.Value == nil || message.Value == "" {
				result = &Message{Type: Failure}
			} else {
				h.mutex.Lock()
				h.storage.Set(message.Key, message.Value)
				h.mutex.Unlock()
				result = &Message{Type: Success}
			}
		case Remove:
			if message.Key == "" || h.storage.Get(message.Key) == nil {
				result = &Message{Type: Failure}
			} else {
				h.mutex.Lock()
				h.storage.Remove(message.Key)
				h.mutex.Unlock()
				result = &Message{Type: Success}
			}
		default:
			log.Println("[INFO] unsupported message", message)
		}
		out <- &messageAndConn{message: result, conn: req.conn}
	}
}
