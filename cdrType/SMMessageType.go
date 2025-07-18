package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	SMMessageTypePresentSubmission       asn.Enumerated = 0
	SMMessageTypePresentDeliveryReport   asn.Enumerated = 1
	SMMessageTypePresentSMServiceRequest asn.Enumerated = 2
	SMMessageTypePresentDelivery         asn.Enumerated = 3
	SMMessageTypePresentT4DeviceTrigger  asn.Enumerated = 4
	SMMessageTypePresentSMDeviceTrigger  asn.Enumerated = 5
)

type SMMessageType struct {
	Value asn.Enumerated
}
