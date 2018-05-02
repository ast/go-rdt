package rdt

import (
	"bytes"
	"fmt"
)

/*
 * Contact
 */
type ContactType uint8

const (
	Group   ContactType = 1
	Private             = 2
	All                 = 3
)

func (c ContactType) String() string {
	switch c {
	case Group:
		return "Group"
	case Private:
		return "Private"
	case All:
		return "All"
	default:
		return fmt.Sprintf("unknown(%d)", c)
	}
}

type Contact struct {
	Deleted bool `json:"-"`
	CallId  uint32
	Type    ContactType
	Name    string
}

func decodeContacts(data []byte) map[int]Contact {
	contacts := make(map[int]Contact)
	r := bytes.NewReader(data)
	contact := make([]byte, 36, 36)
	for i := 0; ; i++ {
		_, err := r.Read(contact)
		if err != nil {
			break
		}
		c := Contact{}

		c.Deleted = contact[4] == 0x00 && contact[5] == 0x00
		c.CallId = uint24ToUint32(contact)
		c.Type = ContactType(contact[3] & 0x03)
		c.Name = bytesToString(contact[4:])

		if !c.Deleted {
			contacts[i] = c
		}
	}
	return contacts
}
