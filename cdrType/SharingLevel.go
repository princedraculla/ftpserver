package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	SharingLevelPresentSHARED    asn.Enumerated = 0
	SharingLevelPresentNONSHARED asn.Enumerated = 1
)

type SharingLevel struct {
	Value asn.Enumerated
}
