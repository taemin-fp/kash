package main

import (
	"io"
	"log"
	"net"
)

type Handler struct {
	storage Storage
}

type messageAndConn struct {
	message *Message
	conn    *Conn
}

func server() {
	handler := Handler{storage: GetStorage()}

	l, err := net.Listen("tcp", ":2934")
	if err != nil {
		log.Panicln(err)
	}
	defer l.Close()

	in := make(chan *messageAndConn)
	out := make(chan *messageAndConn)

	go handler.handler(in, out)
	go handler.responder(out)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("[WARN] cannot accept", conn)
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
			continue
		}
	}
}

func (h *Handler) handler(in <-chan *messageAndConn, out chan<- *messageAndConn) {
	for req := range in {
		message := req.message

		var result *Message
		switch message.Type {
		case Get:
			value := h.storage.Get(message.Key)
			result = &Message{Type: Success, Value: value}
		case Set:
			if message.Value == nil || message.Value == "" {
				result = &Message{Type: Failure}
			} else {
				h.storage.Set(message.Key, message.Value)
				result = &Message{Type: Success}
			}
		case Remove:
			if message.Key == "" || h.storage.Get(message.Key) == nil {
				result = &Message{Type: Failure}
			} else {
				h.storage.Remove(message.Key)
				result = &Message{Type: Success}
			}
		default:
			log.Println("[INFO] unsupported message", message)
		}
		out <- &messageAndConn{message: result, conn: req.conn}
	}
}
