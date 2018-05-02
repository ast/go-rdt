package rdt

import (
	"bytes"
	"encoding/json"
)

/*
 * TextMessage
 */
type TextMessage struct {
	Deleted bool `json:"-"`
	Message string
}

func (tm TextMessage) String() string {
	bs, _ := json.MarshalIndent(tm, "", "  ")
	return string(bs)
}

func decodeTextMessages(data []byte) map[int]TextMessage {
	textMessages := make(map[int]TextMessage)
	r := bytes.NewReader(data)
	message := make([]byte, 288, 288)
	for i := 0; ; i++ {
		_, err := r.Read(message)
		if err != nil {
			break
		}
		tm := TextMessage{}

		tm.Deleted = message[0] == 0x00
		tm.Message = bytesToString(message)
		if !tm.Deleted {
			textMessages[i] = tm
		}
	}
	return textMessages
}
