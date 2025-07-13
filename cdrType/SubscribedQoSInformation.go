package cdrType

type SubscribedQoSInformation struct { /* Sequence Type */
	FiveQi        *int64                       `ber:"tagNum:1,optional"`
	ARP           *AllocationRetentionPriority `ber:"tagNum:2,optional"`
	PriorityLevel *int64                       `ber:"tagNum:3,optional"`
}
