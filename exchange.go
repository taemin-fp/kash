package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"net"
)

type Conn struct {
	conn net.Conn
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{conn: conn}
}

func (c *Conn) Send(message *Message) error {
	buffer, err := serialize(message)
	if err != nil {
		return err
	}
	bufferSize := len(buffer)
	lenBuffer := make([]byte, 4)
	remain := bufferSize

	binary.BigEndian.PutUint32(lenBuffer, uint32(bufferSize))
	_, err = c.conn.Write(lenBuffer)
	if err != nil {
		return err
	}

	for remain > 0 {
		n, err := c.conn.Write(buffer)
		remain -= n
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) Receive() (*Message, error) {
	lenBuffer := make([]byte, 4)
	messageBuffer := make([]byte, 0)
	buffer := make([]byte, 8192)

	_, err := c.conn.Read(lenBuffer)
	if err != nil {
		return nil, err
	}
	messageLen := binary.BigEndian.Uint32(lenBuffer)
	received := uint32(0)

	for received < messageLen {
		n, err := c.conn.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		messageBuffer = append(messageBuffer, buffer[:n]...)
		received += uint32(n)
	}

	message, err := deserialize(messageBuffer)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func serialize(message *Message) ([]byte, error) {
	var codecBuffer bytes.Buffer
	encoder := gob.NewEncoder(&codecBuffer)
	if err := encoder.Encode(*message); err != nil {
		return nil, err
	}

	return codecBuffer.Bytes(), nil
}

func deserialize(buffer []byte) (*Message, error) {
	codecBuffer := bytes.NewBuffer(buffer)
	message := Message{}
	decoder := gob.NewDecoder(codecBuffer)
	if err := decoder.Decode(&message); err != nil {
		return nil, err
	}

	return &message, nil
}
