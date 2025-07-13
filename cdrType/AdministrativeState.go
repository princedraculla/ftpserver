package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	AdministrativeStatePresentLOCKED       asn.Enumerated = 0
	AdministrativeStatePresentUNLOCKED     asn.Enumerated = 1
	AdministrativeStatePresentSHUTTINGDOWN asn.Enumerated = 2
)

type AdministrativeState struct {
	Value asn.Enumerated
}
