package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	DelayToleranceIndicatorPresentDTSupported    asn.Enumerated = 0
	DelayToleranceIndicatorPresentDTNotSupported asn.Enumerated = 1
)

type DelayToleranceIndicator struct {
	Value asn.Enumerated
}
