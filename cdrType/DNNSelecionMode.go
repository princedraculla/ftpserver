package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	DNNSelectionModePresentUEorNetworkProvidedSubscriptionVerified asn.Enumerated = 0
	DNNSelectionModePresentUEProvidedSubscriptionNotVerified       asn.Enumerated = 1
	DNNSelectionModePresentNetworkProvidedSubscriptionNotVerified  asn.Enumerated = 2
)

type DNNSelectionMode struct {
	Value asn.Enumerated
}
