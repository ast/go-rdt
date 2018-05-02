package rdt

import (
	"bytes"
	"encoding/json"
)

/*
 * RXGroupList
 */
type RXGroupList struct {
	Deleted         bool `json:"-"`
	Name            string
	ContactIndicies []uint16
}

func (rx RXGroupList) String() string {
	bs, _ := json.MarshalIndent(rx, "", "  ")
	return string(bs)
}

func decodeRXGroupLists(data []byte) map[int]RXGroupList {
	rxGroupLists := make(map[int]RXGroupList)
	r := bytes.NewReader(data)
	rxGroupList := make([]byte, 96, 96)
	for i := 0; ; i++ {
		_, err := r.Read(rxGroupList)
		if err != nil {
			break
		}
		rx := RXGroupList{}

		rx.Deleted = rxGroupList[0] == 0
		rx.Name = bytesToString(rxGroupList[0:32])
		rx.ContactIndicies = bytesToUint16s(rxGroupList[32:])

		if !rx.Deleted {
			rxGroupLists[i] = rx
		}
	}
	return rxGroupLists
}
