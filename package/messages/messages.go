package messages

import (
	"encoding/json"
)

type Action uint32

const (
	LIST     Action = 0
	RELAY    Action = 1
	RESPONSE Action = 2
)

type Message struct {
	Id     uint64
	Action Action
	Body   []byte
}

func Build(b []byte) (*Message, error) {
	var m Message

	err := json.Unmarshal(b, &m)

	return &m, err
}

func (m *Message) Serialize() ([]byte, error) {
	return json.Marshal(m)
}
