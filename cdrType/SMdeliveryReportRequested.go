package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	SMdeliveryReportRequestedPresentYes asn.Enumerated = 0
	SMdeliveryReportRequestedPresentNo  asn.Enumerated = 1
)

type SMdeliveryReportRequested struct {
	Value asn.Enumerated
}
