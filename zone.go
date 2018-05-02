package rdt

import (
	"bytes"
)

/*
 * Zone
 */
type Zone struct {
	Deleted         bool `json:"-"`
	Name            string
	ChannelIndicies []uint16
}

func decodeZones(data []byte) map[int]Zone {
	zones := make(map[int]Zone)
	r := bytes.NewReader(data)
	zone := make([]byte, 64)
	for i := 0; ; i++ {
		_, err := r.Read(zone)
		if err != nil {
			break
		}
		z := Zone{}
		z.Deleted = zone[0] == 0
		z.Name = bytesToString(zone[0:32])
		z.ChannelIndicies = bytesToUint16s(zone[32:])
		if !z.Deleted {
			zones[i] = z
		}
	}
	return zones
}
