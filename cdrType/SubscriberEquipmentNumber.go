package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type SubscriberEquipmentNumber struct { /* Set Type */
	SubscriberEquipmentNumberType SubscriberEquipmentType `ber:"tagNum:0"`
	SubscriberEquipmentNumberData asn.OctetString         `ber:"tagNum:1"`
}
