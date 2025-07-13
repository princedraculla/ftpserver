package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	MAPDUSteeringFunctionalityPresentMPTCP   asn.Enumerated = 0
	MAPDUSteeringFunctionalityPresentATSSSLL asn.Enumerated = 1
)

type MAPDUSteeringFunctionality struct {
	Value asn.Enumerated
}
