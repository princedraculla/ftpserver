package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	PriorityTypePresentLow    asn.Enumerated = 0
	PriorityTypePresentNormal asn.Enumerated = 1
	PriorityTypePresentHigh   asn.Enumerated = 2
)

type PriorityType struct {
	Value asn.Enumerated
}
