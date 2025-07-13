package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	ManagementOperationStatusPresentOPERATIONSUCCEEDED asn.Enumerated = 0
	ManagementOperationStatusPresentOPERATIONFAILED    asn.Enumerated = 1
)

type ManagementOperationStatus struct {
	Value asn.Enumerated
}
