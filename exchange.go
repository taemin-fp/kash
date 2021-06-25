package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

func serialize(message *Message) []byte {
	var codecBuffer bytes.Buffer
	encoder := gob.NewEncoder(&codecBuffer)
	if err := encoder.Encode(*message); err != nil {
		log.Panicln(err)
	}

	return codecBuffer.Bytes()
}
