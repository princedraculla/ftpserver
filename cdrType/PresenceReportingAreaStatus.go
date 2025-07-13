package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	PresenceReportingAreaStatusPresentInsideArea  asn.Enumerated = 0
	PresenceReportingAreaStatusPresentOutsideArea asn.Enumerated = 1
	PresenceReportingAreaStatusPresentInactive    asn.Enumerated = 2
	PresenceReportingAreaStatusPresentUnknown     asn.Enumerated = 3
)

type PresenceReportingAreaStatus struct {
	Value asn.Enumerated
}
