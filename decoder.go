package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	cdrStorage = "/home/amir/ftp_files/var/CDRs/1"
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

func (c *CDRFile) DecodeCDRFile(filename string) (*CDRFile, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	log.Printf("Decoding file %s: size %d bytes", filename, len(data))

	if len(data) < 54 {
		return nil, fmt.Errorf("file too short")
	}

	cdrFile := &CDRFile{}
	offset := 0

	// Decode CdrFileHeader
	h := &cdrFile.Hdr
	h.FileLength = binary.BigEndian.Uint32(data[offset : offset+4])
	fmt.Printf("the file length:  %d \n", h.FileLength)
	if uint32(len(data)) < h.FileLength {
		return nil, fmt.Errorf("file too short: expected %d bytes, got %d", h.FileLength, len(data))
	}
	offset += 4
	h.HeaderLength = binary.BigEndian.Uint32(data[offset : offset+4])
	fmt.Printf("the Header Length : %d \n", h.HeaderLength)
	if h.HeaderLength < 54 {
		return nil, fmt.Errorf("invalid header length: %d", h.HeaderLength)
	}
	offset += 4
	highIdentifier := data[offset]
	h.HighReleaseIdentifier = highIdentifier >> 5
	h.HighVersionIdentifier = highIdentifier & 0x1F
	offset++
	fmt.Printf("the High Release Identifier : %d \n", h.HighReleaseIdentifier)
	fmt.Printf("the High Version Identifier : %d \n", h.HighVersionIdentifier)
	lowIdentifier := data[offset]
	h.LowReleaseIdentifier = lowIdentifier >> 5
	h.LowVersionIdentifier = lowIdentifier & 0x1F
	offset++
	fmt.Printf("the Low Release Identifier : %d \n", h.LowReleaseIdentifier)
	fmt.Printf("the Low Release Identifier : %d \n", h.LowVersionIdentifier)
	h.FileOpeningTimeStamp = decodeCdrHdrTimeStamp(binary.BigEndian.Uint32(data[offset : offset+4]))
	offset += 4
	fmt.Printf("the File OpeningTimeStamp : %d \n", h.FileOpeningTimeStamp)
	h.TimestampWhenLastCdrWasAppendedToFile = decodeCdrHdrTimeStamp(binary.BigEndian.Uint32(data[offset : offset+4]))
	offset += 4
	fmt.Printf("the Timestamp When Last CDR Was Appended To File : %d \n", h.TimestampWhenLastCdrWasAppendedToFile)
	h.NumberOfCdrsInFile = binary.BigEndian.Uint32(data[offset : offset+4])
	fmt.Printf("the Number Of CDRs In File : %d \n", h.NumberOfCdrsInFile)
	if h.NumberOfCdrsInFile > 100 { // Arbitrary limit to catch corruption
		return nil, fmt.Errorf("invalid NumberOfCdrsInFile: %d", h.NumberOfCdrsInFile)
	}
	offset += 4
	h.FileSequenceNumber = binary.BigEndian.Uint32(data[offset : offset+4])
	fmt.Printf("the File Sequence Number : %d \n", h.FileSequenceNumber)
	offset += 4
	h.FileClosureTriggerReason = FileClosureTriggerReasonType(data[offset])
	fmt.Printf("the File Closure Trigger Reason : %d \n", h.FileClosureTriggerReason)
	offset++
	copy(h.IPAddressOfNodeThatGeneratedFile[:], data[offset:offset+20])
	offset += 20
	fmt.Printf("the IP Address Of Node That Generated File : %d \n", h.IPAddressOfNodeThatGeneratedFile[:])
	h.LostCdrIndicator = data[offset]
	offset++
	h.LengthOfCdrRoutingFilter = binary.BigEndian.Uint16(data[offset : offset+2])
	fmt.Printf("the Length Of CDR Routing Filter : %d \n", h.LengthOfCdrRoutingFilter)
	if offset+int(h.LengthOfCdrRoutingFilter) > len(data) {
		return nil, fmt.Errorf("file too short for CDRRoutingFiler: need %d bytes, have %d", offset+int(h.LengthOfCdrRoutingFilter), len(data))
	}
	offset += 2
	h.CDRRoutingFiler = make([]byte, h.LengthOfCdrRoutingFilter)
	copy(h.CDRRoutingFiler, data[offset:offset+int(h.LengthOfCdrRoutingFilter)])
	offset += int(h.LengthOfCdrRoutingFilter)
	fmt.Printf("the CDR Routing Filter : %d \n", h.CDRRoutingFiler[:])
	h.LengthOfPrivateExtension = binary.BigEndian.Uint16(data[offset : offset+2])
	fmt.Printf("the Length oF Private Extension : %d \n", h.LengthOfPrivateExtension)
	if offset+int(h.LengthOfPrivateExtension) > len(data) {
		return nil, fmt.Errorf("file too short for PrivateExtension: need %d bytes, have %d", offset+int(h.LengthOfPrivateExtension), len(data))
	}
	offset += 2
	h.PrivateExtension = make([]byte, h.LengthOfPrivateExtension)
	copy(h.PrivateExtension, data[offset:offset+int(h.LengthOfPrivateExtension)])
	offset += int(h.LengthOfPrivateExtension)
	fmt.Printf("the Private Extension : %d \n", h.PrivateExtension[:])
	h.HighReleaseIdentifierExtension = data[offset]
	offset++
	fmt.Printf("the High Release Identifier Extension : %d \n", h.HighReleaseIdentifierExtension)
	h.LowReleaseIdentifierExtension = data[offset]
	offset++
	fmt.Printf("the Low Release Identifier Extension : %d \n", h.LowReleaseIdentifierExtension)

	if uint32(offset) != h.HeaderLength {
		return nil, fmt.Errorf("header length mismatch: expected %d, processed %d", h.HeaderLength, offset)
	}

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
			return nil, fmt.Errorf("file too short for CDR %d data: need %d bytes, have %d", i, offset+int(cdr.Hdr.CdrLength), len(data))
		}
		cdr.CdrByte = make([]byte, cdr.Hdr.CdrLength)
		copy(cdr.CdrByte, data[offset:offset+int(cdr.Hdr.CdrLength)])
		offset += int(cdr.Hdr.CdrLength)
		cdrFile.CdrList = append(cdrFile.CdrList, cdr)
		fmt.Printf("the CDR List : %d \n", cdrFile.CdrList[:])
	}
	if uint32(len(data)) != h.FileLength {
		log.Printf("Warning: File length mismatch, expected %d, got %d", h.FileLength, len(data))
	}

	return cdrFile, nil
}

func CDRFiles() ([]string, error) {
	entries, err := os.ReadDir(cdrStorage)
	if err != nil {
		return nil, err
	}
	var files []string

	for _, entrie := range entries {
		if !entrie.IsDir() {
			files = append(files, filepath.Join(cdrStorage, entrie.Name()))
		}
	}
	fmt.Println("founded files : ", files)

	return files, nil
}
