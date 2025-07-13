package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	AccessTypePresentThreeGPPAccess    asn.Enumerated = 0
	AccessTypePresentNonThreeGPPAccess asn.Enumerated = 1
)

type AccessType struct {
	Value asn.Enumerated
}
