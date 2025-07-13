package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	SmsIndicationPresentSMSSupported    asn.Enumerated = 0
	SmsIndicationPresentSMSNotSupported asn.Enumerated = 1
)

type SmsIndication struct {
	Value asn.Enumerated
}
