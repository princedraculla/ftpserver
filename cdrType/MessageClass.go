package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	MessageClassPresentPersonal           asn.Enumerated = 0
	MessageClassPresentAdvertisement      asn.Enumerated = 1
	MessageClassPresentInformationService asn.Enumerated = 2
	MessageClassPresentAuto               asn.Enumerated = 3
)

type MessageClass struct {
	Value asn.Enumerated
}
