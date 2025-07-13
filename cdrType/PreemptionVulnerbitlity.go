package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	PreemptionVulnerabilityPresentNOTPREEMPTABLE asn.Enumerated = 0
	PreemptionVulnerabilityPresentPREEMPTABLE    asn.Enumerated = 1
)

type PreemptionVulnerability struct {
	Value asn.Enumerated
}
