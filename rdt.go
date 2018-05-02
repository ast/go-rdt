package rdt

import (
	"encoding/json"
)

type RDT struct {
	General      General
	TextMessages map[int]TextMessage
	RXGroupList  map[int]RXGroupList
	ScanList     map[int]ScanList
	Channels     map[int]Channel
	Contacts     map[int]Contact
	Zones        map[int]Zone
}

func (r RDT) String() string {
	bs, _ := json.MarshalIndent(r, "", "  ")
	return string(bs)
}

func Decode(data []byte) *RDT {

	rdt := &RDT{}

	rdt.RXGroupList = decodeRXGroupLists(data[0x0ee45:0x14c05])
	rdt.Zones = decodeZones(data[0x14c05:0x18a85])
	rdt.ScanList = decodeScanLists(data[0x18a85:0x1f015])
	rdt.Channels = decodeChannels(data[0x1f025:0x2ea25])
	rdt.General = decodeGeneral(data[0x2265:0x22F5])
	rdt.TextMessages = decodeTextMessages(data[0x23a5:0x5be5])
	rdt.Contacts = decodeContacts(data[0x061a5:0x0ee45])

	return rdt
}
