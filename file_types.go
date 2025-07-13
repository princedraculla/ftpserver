package main

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