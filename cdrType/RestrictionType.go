package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	RestrictionTypePresentAllowedAreas    asn.Enumerated = 0
	RestrictionTypePresentNotAllowedAreas asn.Enumerated = 1
)

type RestrictionType struct {
	Value asn.Enumerated
}
