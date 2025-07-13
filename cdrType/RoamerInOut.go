package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	RoamerInOutPresentRoamerInBound  asn.Enumerated = 0
	RoamerInOutPresentRoamerOutBound asn.Enumerated = 1
)

type RoamerInOut struct {
	Value asn.Enumerated
}
