package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
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

func serialize(message *Message) (result []byte, err error) {
	switch message.Type {
	case Get:
		result = []byte(Get + " " + message.Key)
	case Set:
		result = []byte(fmt.Sprintf("%s %s %s", Set, message.Key, message.Value))
	case Remove:
		result = []byte(Remove + " " + message.Key)
	case Success:
		result = []byte(fmt.Sprintf("%s %s", Success, message.Value))
	case Failure:
		result = []byte(fmt.Sprintf("%s %s", Failure, message.Value))
	default:
		result = nil
		err = errors.New("no matching message type of " + message.Type)
	}
	return
}

func deserialize(buffer []byte) (message *Message, err error) {
	command := strings.Split(string(buffer), " ")
	commandLen := len(command)
	if commandLen < 1 {
		return nil, errors.New("don't know how to handle " + string(buffer))
	}

	switch command[0] {
	case Get:
		if commandLen == 2 {
			message = &Message{
				Type: Get,
				Key:  command[1],
			}
		} else {
			err = errors.New("illegal get " + string(buffer))
		}
	case Set:
		if commandLen == 3 {
			message = &Message{
				Type:  Set,
				Key:   command[1],
				Value: command[2],
			}
		} else {
			err = errors.New("illegal set " + string(buffer))
		}
	case Remove:
		if commandLen == 2 {
			message = &Message{
				Type: Remove,
				Key:  command[1],
			}
		} else {
			err = errors.New("illegal remove " + string(buffer))
		}
	case Success:
		if commandLen == 2 {
			message = &Message{
				Type:  Success,
				Value: command[1],
			}
		} else if commandLen == 1 {
			message = &Message{
				Type: Success,
			}
		} else {
			err = errors.New("illegal success " + string(buffer))
		}
	case Failure:
		message = &Message{
			Type: Failure,
		}
	default:
		err = errors.New("don't know how to handle " + command[0])
	}
	return
}
