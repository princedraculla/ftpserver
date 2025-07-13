package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	CoreNetworkTypePresentFiveGC asn.Enumerated = 0
	CoreNetworkTypePresentEPC    asn.Enumerated = 1
)

type CoreNetworkType struct {
	Value asn.Enumerated
}
