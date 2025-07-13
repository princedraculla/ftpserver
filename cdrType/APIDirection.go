package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	APIDirectionPresentInvocation   asn.Enumerated = 0
	APIDirectionPresentNotification asn.Enumerated = 1
)

type APIDirection struct {
	Value asn.Enumerated
}
