package rdt

import (
	"bytes"
)

/*
 * ScanList
 */
type ScanList struct {
	Deleted         bool `json:"-"`
	Name            string
	ChannelIndicies []uint16
}

func decodeScanLists(data []byte) map[int]ScanList {
	scanLists := make(map[int]ScanList)
	r := bytes.NewReader(data)
	// scanlist is 104 bytes
	scanList := make([]byte, 104, 104)
	for i := 0; ; i++ {
		_, err := r.Read(scanList)
		if err != nil {
			break
		}
		s := ScanList{}

		s.Deleted = scanList[0] == 0x00
		s.Name = bytesToString(scanList[0:32])
		s.ChannelIndicies = bytesToUint16s(scanList[42:])
		if !s.Deleted {
			scanLists[i] = s
		}
	}
	return scanLists
}
