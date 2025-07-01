package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type TsNumberIdentifier uint8

const (
	TS32005 TsNumberIdentifier = 0
	TS32015 TsNumberIdentifier = 1
	TS32205 TsNumberIdentifier = 2
	TS32215 TsNumberIdentifier = 3
	TS32225 TsNumberIdentifier = 4
	TS32235 TsNumberIdentifier = 5
	TS32250 TsNumberIdentifier = 6
	TS32251 TsNumberIdentifier = 7
	TS32260 TsNumberIdentifier = 9
	TS32270 TsNumberIdentifier = 10
	TS32271 TsNumberIdentifier = 11
	TS32272 TsNumberIdentifier = 12
	TS32273 TsNumberIdentifier = 13
	TS32275 TsNumberIdentifier = 14
	TS32274 TsNumberIdentifier = 15
	TS32277 TsNumberIdentifier = 16
	TS32296 TsNumberIdentifier = 17
	TS32278 TsNumberIdentifier = 18
	TS32253 TsNumberIdentifier = 19
	TS32255 TsNumberIdentifier = 20
	TS32254 TsNumberIdentifier = 21
	TS32256 TsNumberIdentifier = 22
	TS28201 TsNumberIdentifier = 23
	TS28202 TsNumberIdentifier = 24
)

type DataRecordFormatType uint8

const (
	BasicEncodingRules DataRecordFormatType = iota + 1
	UnalignedPackedEncodingRules
	AlignedPackedEncodingRules1
	XMLEncodingRules
)

type CDRFile struct {
	Hdr     CdrFileHeader
	CdrList []CDR
}
type CDR struct {
	Hdr     CDRHeader
	CdrByte []byte
}

type CdrFileHeader struct {
	FileLength                            uint32
	HeaderLength                          uint32
	HighReleaseIdentifier                 uint8 // octet 9 bit 6..8
	HighVersionIdentifier                 uint8 // octet 9 bit 1..5
	LowReleaseIdentifier                  uint8 // octet 10 bit 6..8
	LowVersionIdentifier                  uint8 // octet 10 bit 1..5
	FileOpeningTimeStamp                  CdrHdrTimeStamp
	TimestampWhenLastCdrWasAppendedToFile CdrHdrTimeStamp
	NumberOfCdrsInFile                    uint32
	FileSequenceNumber                    uint32
	FileClosureTriggerReason              FileClosureTriggerReasonType
	IPAddressOfNodeThatGeneratedFile      [20]byte // ip address in ipv6 format
	LostCdrIndicator                      uint8
	LengthOfCdrRoutingFilter              uint16
	CDRRoutingFiler                       []byte //vendor specific
	LengthOfPrivateExtension              uint16
	PrivateExtension                      []byte // vendor specific
	HighReleaseIdentifierExtension        uint8
	LowReleaseIdentifierExtension         uint8
}

type CDRHeader struct {
	CdrLength                  uint16
	ReleaseIdentifier          ReleaseIdentifierType // octect 3 bit 6..8
	VersionIdentifier          uint8
	DataRecordFormat           DataRecordFormatType
	TsNumber                   TsNumberIdentifier
	ReleaseIdentifierExtension uint8
}
type CdrHdrTimeStamp struct {
	MonthLocal  uint8
	DateLocal   uint8
	HourLocal   uint8
	MinuteLocal uint8

	// bit set to "1" expresses "+" or bit set to "0" expresses "-" time deviation
	SignOfTheLocalTimeDifferentialFromUtc uint8
	HourDeviation                         uint8
	MinuteDeviation                       uint8
}

type FileClosureTriggerReasonType uint8
type ReleaseIdentifierType uint8

func decodeCdrHdrTimeStamp(ts uint32) CdrHdrTimeStamp {
	return CdrHdrTimeStamp{
		MonthLocal:                            uint8(ts >> 28),
		DateLocal:                             uint8((ts >> 23) & 0x1F),
		HourLocal:                             uint8((ts >> 18) & 0x1F),
		MinuteLocal:                           uint8((ts >> 12) & 0x3F),
		SignOfTheLocalTimeDifferentialFromUtc: uint8((ts >> 11) & 0x1),
		HourDeviation:                         uint8((ts >> 6) & 0x1F),
		MinuteDeviation:                       uint8(ts & 0x3F),
	}
}

func DecodeCDRFile(filename string) (*CDRFile, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(data) < 54 {
		return nil, fmt.Errorf("file too short")
	}

	cdrFile := &CDRFile{}
	offset := 0

	// Decode CdrFileHeader
	h := &cdrFile.Hdr
	h.FileLength = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4
	h.HeaderLength = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4
	highIdentifier := data[offset]
	h.HighReleaseIdentifier = highIdentifier >> 5
	h.HighVersionIdentifier = highIdentifier & 0x1F
	offset++
	lowIdentifier := data[offset]
	h.LowReleaseIdentifier = lowIdentifier >> 5
	h.LowVersionIdentifier = lowIdentifier & 0x1F
	offset++
	h.FileOpeningTimeStamp = decodeCdrHdrTimeStamp(binary.BigEndian.Uint32(data[offset : offset+4]))
	offset += 4
	h.TimestampWhenLastCdrWasAppendedToFile = decodeCdrHdrTimeStamp(binary.BigEndian.Uint32(data[offset : offset+4]))
	offset += 4
	h.NumberOfCdrsInFile = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4
	h.FileSequenceNumber = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4
	h.FileClosureTriggerReason = FileClosureTriggerReasonType(data[offset])
	offset++
	copy(h.IPAddressOfNodeThatGeneratedFile[:], data[offset:offset+20])
	offset += 20
	h.LostCdrIndicator = data[offset]
	offset++
	h.LengthOfCdrRoutingFilter = binary.BigEndian.Uint16(data[offset : offset+2])
	offset += 2
	h.CDRRoutingFiler = make([]byte, h.LengthOfCdrRoutingFilter)
	copy(h.CDRRoutingFiler, data[offset:offset+int(h.LengthOfCdrRoutingFilter)])
	offset += int(h.LengthOfCdrRoutingFilter)
	h.LengthOfPrivateExtension = binary.BigEndian.Uint16(data[offset : offset+2])
	offset += 2
	h.PrivateExtension = make([]byte, h.LengthOfPrivateExtension)
	copy(h.PrivateExtension, data[offset:offset+int(h.LengthOfPrivateExtension)])
	offset += int(h.LengthOfPrivateExtension)
	h.HighReleaseIdentifierExtension = data[offset]
	offset++
	h.LowReleaseIdentifierExtension = data[offset]
	offset++

	// Decode CDRs
	for i := uint32(0); i < h.NumberOfCdrsInFile; i++ {
		if offset+5 > len(data) {
			return nil, fmt.Errorf("file too short for CDR %d", i)
		}
		var cdr CDR
		cdr.Hdr.CdrLength = binary.BigEndian.Uint16(data[offset : offset+2])
		offset += 2
		identifier := data[offset]
		cdr.Hdr.ReleaseIdentifier = ReleaseIdentifierType(identifier >> 5)
		cdr.Hdr.VersionIdentifier = identifier & 0x1F
		offset++
		oct4 := data[offset]
		cdr.Hdr.DataRecordFormat = DataRecordFormatType(oct4 >> 5)
		cdr.Hdr.TsNumber = TsNumberIdentifier(oct4 & 0x1F)
		offset++
		cdr.Hdr.ReleaseIdentifierExtension = data[offset]
		offset++
		if offset+int(cdr.Hdr.CdrLength) > len(data) {
			return nil, fmt.Errorf("file too short for CDR %d data", i)
		}
		cdr.CdrByte = make([]byte, cdr.Hdr.CdrLength)
		copy(cdr.CdrByte, data[offset:offset+int(cdr.Hdr.CdrLength)])
		offset += int(cdr.Hdr.CdrLength)
		cdrFile.CdrList = append(cdrFile.CdrList, cdr)
	}

	return cdrFile, nil
}
