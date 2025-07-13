package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	SMReplyPathRequestedPresentNoReplyPathSet asn.Enumerated = 0
	SMReplyPathRequestedPresentReplyPathSet   asn.Enumerated = 1
)

type SMReplyPathRequested struct {
	Value asn.Enumerated
}
