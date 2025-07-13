package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	RegistrationMessageTypePresentInitial        asn.Enumerated = 0
	RegistrationMessageTypePresentMobility       asn.Enumerated = 1
	RegistrationMessageTypePresentPeriodic       asn.Enumerated = 2
	RegistrationMessageTypePresentEmergency      asn.Enumerated = 3
	RegistrationMessageTypePresentDeregistration asn.Enumerated = 4
)

type RegistrationMessageType struct {
	Value asn.Enumerated
}
