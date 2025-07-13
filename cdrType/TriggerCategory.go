package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	TriggerCategoryPresentImmediateReport asn.Enumerated = 0
	TriggerCategoryPresentDeferredReport  asn.Enumerated = 1
)

type TriggerCategory struct {
	Value asn.Enumerated
}
