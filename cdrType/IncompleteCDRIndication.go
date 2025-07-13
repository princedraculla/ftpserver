package cdrType

type IncompleteCDRIndication struct { /* Sequence Type */
	InitialLost     *bool `ber:"tagNum:0,optional"`
	UpdateLost      *bool `ber:"tagNum:1,optional"`
	TerminationLost *bool `ber:"tagNum:2,optional"`
}
