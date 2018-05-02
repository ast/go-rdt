package rdt

import (
	"encoding/json"
)

/*
 * General
 */
type General struct {
	InfoScreenLine1        string // Unicode	0	160	Maximum length: 10 characters.
	InfoScreenLine2        string // Unicode	160	160	Maximum length: 10 characters.
	MonitorType            bool   // Binary	515	1	0=silent, 1=open
	DisableAllLeds         bool   // Binary	517	1	0=off, 1=on
	TalkPermitTone         uint8  // Binary	520	2	0=none, 1=digital, 2=analog, 3=both
	PasswordAndLockEnable  bool   // Binary	522	1	1=off, 0=on
	CHFreeIndicationTone   bool   // Binary	523	1	1=off, 0=on
	DisableAllTone         bool   // Binary	525	1	1=off, 0=on
	SaveModeReceive        bool   // Binary	526	1	0=off, 1=on
	SavePreamble           bool   // Binary	527	1	0=off, 1=on
	IntroScreen            bool   // Binary 531	1	0=charstring, 1=picture
	RadioId                uint32 // Binary 544	24	max value is 16776415
	TxPreamble             uint8  // Binary 576	8	ms=N*60, where N<=0<=144
	GroupCallHangTime      uint8  // Binary	584	8	time in ms, ms=N*100, N<=70, N must be multiple of 5
	PrivateCallHangTime    uint8  // Binary	592	8	time in ms, ms=N*100, N<=70, N must be multiple of 5
	VoxSensitivity         uint8  // Binary	600	8
	RxLowBatteryInterval   uint8  // Binary	624	8	time in seconds, s=N*5, N<=127
	CallAlertTone          uint8  // Binary	632	8	0=Continue, otherwise time in seconds, s=N*5, N<=240
	LoneWorkerRespTime     uint8  // Binary	640	8
	LoneWorkerReminderTime uint8  // Binary	648	8
	ScanDigitalHangTime    uint8  // Binary	664	8	time in ms, ms=N*5, 5<=N<=100; default N=10
	ScanAnalogHangTime     uint8  // Binary	672	8	time in ms, ms=N*5, 5<=N<=100; default N=10
	Unknown1               uint8  // Binary	680	8	meaning still unknown, do not edit
	KeypadLockTime         uint8  // Binary	688	8	1=5s, 2=10s, 3=15s, 255=manual
	Mode                   uint8  // Binary	696	8	0=mr, 255=ch
	PowerOnPassword        uint32 // RevBCD	704	32	8 digits
	RadioProgPassword      uint32 // RevBCD	736	32	8 digits
	PcProgPassword         string // Ascii	768	64	Converted to lower case; if not set, set to all 0xFF.
	RadioName              string // Unicode	896	256	Maximum length: 16 characters.
}

func (g General) String() string {
	bs, _ := json.MarshalIndent(g, "", "  ")
	return string(bs)
}

func decodeGeneral(data []byte) General {
	g := General{}

	// Info
	g.InfoScreenLine1 = bytesToString(data[0:20])
	g.InfoScreenLine2 = bytesToString(data[20:40])
	// Byte 64
	g.MonitorType = bitToBool(data[64], 4)
	g.DisableAllLeds = bitToBool(data[64], 2)
	// Byte 65
	g.TalkPermitTone = (data[65] & 0xc0) >> 6
	g.PasswordAndLockEnable = bitToBool(data[65], 5)
	g.CHFreeIndicationTone = bitToBool(data[65], 4)
	g.DisableAllTone = bitToBool(data[65], 2)
	g.SaveModeReceive = bitToBool(data[65], 1)
	g.SavePreamble = bitToBool(data[65], 0)
	g.IntroScreen = true

	// Bytes 72-83
	g.TxPreamble = data[72]
	g.GroupCallHangTime = data[72]
	g.PrivateCallHangTime = data[73]
	g.VoxSensitivity = data[74]
	g.RxLowBatteryInterval = data[75]
	g.CallAlertTone = data[76]
	g.LoneWorkerRespTime = data[77]
	g.LoneWorkerReminderTime = data[78]
	g.ScanDigitalHangTime = data[79]
	g.ScanAnalogHangTime = data[80]
	g.Unknown1 = data[81]
	g.KeypadLockTime = data[82]
	g.Mode = data[83]

	//g.PowerOnPassword = bcdDecodeForFrequency(data[88:92])
	//g.RadioProgPassword = bcdDecodeForFrequency(data[92:96])
	// ASCII string
	g.PcProgPassword = string(data[96:104])

	g.RadioId = uint24ToUint32(data[68:])

	g.RadioName = bytesToString(data[112:])

	return g
}
