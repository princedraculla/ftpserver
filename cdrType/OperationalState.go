package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	OperationalStatePresentENABLED  asn.Enumerated = 0
	OperationalStatePresentDISABLED asn.Enumerated = 1
)

type OperationalState struct {
	Value asn.Enumerated
}
