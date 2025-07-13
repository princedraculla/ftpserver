package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	PreemptionCapabilityPresentNOTPREEMPT asn.Enumerated = 0
	PreemptionCapabilityPresentMAYPREEMPT asn.Enumerated = 1
)

type PreemptionCapability struct {
	Value asn.Enumerated
}
