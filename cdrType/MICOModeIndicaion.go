package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	MICOModeIndicationPresentMICOMode   asn.Enumerated = 0
	MICOModeIndicationPresentNoMICOMode asn.Enumerated = 1
)

type MICOModeIndication struct {
	Value asn.Enumerated
}
