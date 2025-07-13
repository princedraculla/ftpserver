package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	ManagementOperationPresentCreateMOI           asn.Enumerated = 0
	ManagementOperationPresentModifyMOIAttributes asn.Enumerated = 1
	ManagementOperationPresentDeleteMOI           asn.Enumerated = 2
)

type ManagementOperation struct {
	Value asn.Enumerated
}
