package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	SteerModeValuePresentActiveStandby asn.Enumerated = 0
	SteerModeValuePresentLoadBalancing asn.Enumerated = 1
	SteerModeValuePresentSmallestDelay asn.Enumerated = 2
	SteerModeValuePresentPriorityBased asn.Enumerated = 3
)

type SteerModeValue struct {
	Value asn.Enumerated
}
