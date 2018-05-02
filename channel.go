package rdt

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/*
 * Channel
 */
type CtcssDcsType uint8

const (
	CTCSS CtcssDcsType = 0x00
	DCS_N CtcssDcsType = 0x02
	DCS_I CtcssDcsType = 0x03
)

type FrequencyType uint32

func (f FrequencyType) String() string {
	// TODO: improve this
	str := fmt.Sprintf("%d", f)
	return fmt.Sprintf("%s.%s", str[0:3], str[3:8])
}

type BandwidthType uint8

const (
	Bandwidth125 BandwidthType = 0
	Bandwidth250               = 1
)

func (b BandwidthType) String() string {
	switch b {
	case Bandwidth125:
		return "12.5"
	case Bandwidth250:
		return "25.0"
	default:
		return fmt.Sprintf("unknown(%d)", b)
	}
}

type PowerType uint8

const (
	Low  PowerType = 0
	High           = 1
)

func (p PowerType) String() string {
	switch p {
	case High:
		return "High"
	case Low:
		return "Low"
	default:
		return fmt.Sprintf("unknown(%d)", p)
	}
}

type Channel struct {
	Deleted           bool `json:"-"`
	LoneWorker        bool
	Squelch           bool
	Autoscan          bool
	Bandwidth         BandwidthType
	ChannelMode       uint8
	ColorCode         uint8
	RepeaterSlot      uint8
	RxOnly            bool
	AllowTalkaround   bool
	DataCallConf      bool
	PrivateCallConf   bool
	Privacy           uint8
	PrivacyNo         uint8
	DisplayPttId      bool
	CompressedUdpHdr  bool
	EmergencyAlarmAck bool
	RxRefFrequency    uint8
	AdmintCriteria    uint8
	Power             PowerType
	Vox               bool
	QtReverse         bool
	ReverseBurst      bool
	TxRefFrequency    uint8
	ContactIndex      uint16
	Tot               uint8
	TotRekeyDelay     uint8
	EmergencySystem   uint8
	ScanListIndex     uint8
	RXGroupListIndex  uint8
	Decode18          uint8
	RxFrequency       FrequencyType
	TxFrequency       FrequencyType
	CtcssDcsDecode    uint16
	CtcssDcsEncode    uint16
	TxSignalingSyst   uint8
	Name              string
}

func (c Channel) String() string {
	bs, _ := json.MarshalIndent(c, "", "  ")
	return string(bs)
}

func decodeChannels(data []byte) map[int]Channel {
	channels := make(map[int]Channel)
	r := bytes.NewReader(data)
	// channel is 64 bytes
	channel := make([]byte, 64, 64)
	//for i := 0; i < 50; i++ {
	for i := 0; ; i++ {
		_, err := r.Read(channel)
		if err != nil {
			break
		}

		// Parse channel
		c := Channel{}
		c.Deleted = channel[16] == 0x00
		// Byte 0
		c.LoneWorker = bitToBool(channel[0], 7)
		c.Squelch = bitToBool(channel[0], 5)
		c.Autoscan = bitToBool(channel[0], 4)
		c.Bandwidth = BandwidthType(bitToUint8(channel[0], 3))

		c.ChannelMode = (channel[0] & 0x03)
		// Byte 1
		c.ColorCode = channel[1] >> 4
		c.RepeaterSlot = (channel[1] & 0x0c) >> 2
		c.RxOnly = bitToBool(channel[1], 1)
		c.AllowTalkaround = bitToBool(channel[1], 0)
		// Byte 2
		c.DataCallConf = bitToBool(channel[2], 7)
		c.PrivateCallConf = bitToBool(channel[2], 6)
		c.Privacy = (channel[2] & 0x30) >> 4
		c.PrivacyNo = channel[2] & 0x0f
		// Byte 3
		c.DisplayPttId = bitToBool(channel[3], 7)
		c.CompressedUdpHdr = bitToBool(channel[3], 6)
		c.EmergencyAlarmAck = bitToBool(channel[3], 3)
		c.RxRefFrequency = channel[3] & 0x03
		// Byte 4
		c.AdmintCriteria = (channel[4] & 0xc0) >> 6
		c.Power = PowerType(bitToUint8(channel[4], 5))
		c.Vox = bitToBool(channel[4], 4)
		c.QtReverse = bitToBool(channel[4], 3)
		c.ReverseBurst = bitToBool(channel[4], 2)
		c.TxRefFrequency = channel[4] & 0x03
		// Byte 5 = empty?
		// Byte 6
		c.ContactIndex = bytesToUint16(channel[6:])
		// Byte 8
		c.Tot = channel[8] & 0x3f
		// Byte 9
		c.TotRekeyDelay = channel[9]
		// Byte 10
		c.EmergencySystem = channel[10] & 0x3f
		// Byte 11
		c.ScanListIndex = channel[11]
		// Byte 12
		c.RXGroupListIndex = channel[12]
		// Byte 14
		c.Decode18 = channel[14]
		// Byte 16
		c.RxFrequency = bcdToFrequency(channel[16:])
		// Byte 20
		c.TxFrequency = bcdToFrequency(channel[20:])
		// Byte 24
		c.CtcssDcsDecode = bytesToUint16(channel[24:])
		// Byte 26
		c.CtcssDcsEncode = bytesToUint16(channel[26:])
		// Byte 29
		c.TxSignalingSyst = channel[29] & 0x03
		// c.RxSignalingSyst
		c.Name = bytesToString(channel[32:])

		if !c.Deleted {
			channels[i] = c
		}
	}
	return channels
}
