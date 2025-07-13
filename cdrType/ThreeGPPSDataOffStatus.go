package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	ThreeGPPPSDataOffStatusPresentActive   asn.Enumerated = 0
	ThreeGPPPSDataOffStatusPresentInactive asn.Enumerated = 1
)

type ThreeGPPPSDataOffStatus struct {
	Value asn.Enumerated
}
