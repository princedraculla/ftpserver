package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	SubscriberEquipmentTypePresentIMEISV        asn.Enumerated = 0
	SubscriberEquipmentTypePresentMAC           asn.Enumerated = 1
	SubscriberEquipmentTypePresentEUI64         asn.Enumerated = 2
	SubscriberEquipmentTypePresentModifiedEUI64 asn.Enumerated = 3
)

type SubscriberEquipmentType struct {
	Value asn.Enumerated
}
