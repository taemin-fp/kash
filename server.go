package main

import (
	"bytes"
	"encoding/gob"
	"io"
	"log"
	"net"
)

type Handler struct {
	storage Storage
}

type messageAndConn struct {
	message *Message
	conn    net.Conn
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

		go readMessage(conn, in)
	}
}

func readMessage(conn net.Conn, in chan *messageAndConn) {
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

	in <- &messageAndConn{message: &message, conn: conn}
}

func (h *Handler) responder(out <-chan *messageAndConn) {
	for res := range out {
		buffer := serialize(res.message)

		n, err := res.conn.Write(buffer)
		if n != len(buffer) {
			log.Printf("[WARN] cannot send response completely; %d bytes sent", n)
		} else if err != nil {
			log.Println("[WARN] cannot send response")
		}
	}
}

func (h *Handler) handler(in <-chan *messageAndConn, out chan *messageAndConn) {
	for req := range in {
		message := req.message

		var result *Message
		switch message.Type {
		case Get:
			value := h.storage.Get(message.Key)
			result = &Message{Type: Success, Value: value}
		case Set:
			log.Printf("[INFO] key %s, value %s", message.Key, message.Value)
			if message.Value == nil || message.Value == "" {
				result = &Message{Type: Failure}
			} else {
				h.storage.Set(message.Key, message.Value)
				result = &Message{Type: Success}
			}
		default:
			log.Println("[INFO] unsupported message", message)
		}
		out <- &messageAndConn{message: result, conn: req.conn}
	}
}
