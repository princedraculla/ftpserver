package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/princedraculla/ftpservertest/asn"
	"github.com/princedraculla/ftpservertest/cdrType"
)

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
		var records cdrType.ChargingRecord
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
		if err := asn.UnmarshalWithParams(cdr.CdrByte, &records, "explicit,choice"); err != nil {
			return nil, fmt.Errorf("error while unmarshaling cdrs : %v ", err)
		}
		jsonData, err := json.MarshalIndent(records, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("marshaling cdrs to json error : %v", err)
		}

		fmt.Println("Unmarshaled CDRs : \n", string(jsonData))
		//fmt.Println("unmarshaled CDRs : ", records)
		spew.Dump("unmarshaled CDRs : ", records, "\n")

	}
	if uint32(len(data)) != h.FileLength {
		log.Printf("Warning: File length mismatch, expected %d, got %d", h.FileLength, len(data))
	}

	return cdrFile, nil
}
