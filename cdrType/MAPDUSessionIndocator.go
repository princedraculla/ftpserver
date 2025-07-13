package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	MAPDUSessionIndicatorPresentMAPDURequest               asn.Enumerated = 0
	MAPDUSessionIndicatorPresentMAPDUNetworkUpgradeAllowed asn.Enumerated = 1
)

type MAPDUSessionIndicator struct {
	Value asn.Enumerated
}
