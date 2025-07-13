package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	LineTypePresentDSL asn.Enumerated = 0
	LineTypePresentPON asn.Enumerated = 1
)

type LineType struct {
	Value asn.Enumerated
}
