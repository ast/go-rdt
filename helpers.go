package rdt

import (
	"bytes"
	"encoding/binary"
	"log"
	"unicode/utf16"
)

/*
 * Decoding helpers
 */

func bitToUint8(b byte, bit uint8) uint8 {
	return (b & (1 << bit)) >> bit
}

func bitToBool(b byte, bit uint8) bool {
	return bitToUint8(b, bit) == 0x01
}

/* Convert uint24 to uint32 */
func uint24ToUint32(bs []byte) uint32 {
	return uint32(bs[2])<<16 | uint32(bs[1])<<8 | uint32(bs[0])
}

func bcdtDecode(bs []byte) {

}

/* 4 bit BCD to uint32 */
func bcdToFrequency(bs []byte) FrequencyType {
	freq := uint32(0)
	freq += uint32(bs[0]&0x0f) * 10
	freq += uint32(bs[0]>>4) * 100
	freq += uint32(bs[1]&0x0f) * 1e3
	freq += uint32(bs[1]>>4) * 10e3
	freq += uint32(bs[2]&0x0f) * 100e3
	freq += uint32(bs[2]>>4) * 1e6
	freq += uint32(bs[3]&0x0f) * 10e6
	freq += uint32(bs[3]>>4) * 100e6
	return FrequencyType(freq)
}

/* reverse 4 bit BCD to uint32 */
func revBCDDecode(bs []byte) uint32 {
	v := uint32(0)
	v += uint32(bs[3]&0x0f) * 10
	v += uint32(bs[3]>>4) * 100
	v += uint32(bs[2]&0x0f) * 1e3
	v += uint32(bs[2]>>4) * 10e3
	v += uint32(bs[1]&0x0f) * 100e3
	v += uint32(bs[1]>>4) * 1e6
	v += uint32(bs[0]&0x0f) * 10e6
	v += uint32(bs[0]>>4) * 100e6
	return v
}

/* Bytes to array of uint16 */
func bytesToUint16s(bs []byte) []uint16 {
	count := len(bs) / 2
	uints := make([]uint16, count, count)
	reader := bytes.NewReader(bs)
	if err := binary.Read(reader, binary.LittleEndian, &uints); err != nil {
		log.Fatal(err)
	}
	/* Return up until first 0x00 byte */
	for i, v := range uints {
		if v == 0 {
			return uints[:i]
		}
	}
	return uints
}

/* Bytes uint16 LE */
func bytesToUint16(bs []byte) uint16 {
	v := binary.LittleEndian.Uint16(bs)
	return v
}

/* UTF-16 bytes to string */
func bytesToString(bs []byte) string {
	utf := bytesToUint16s(bs)
	return string(utf16.Decode(utf))
}
