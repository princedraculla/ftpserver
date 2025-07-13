package cdrType

type SingleNSSAI struct { /* Sequence Type */
	SST SliceServiceType     `ber:"tagNum:0"`
	SD  *SliceDifferentiator `ber:"tagNum:1,optional"`
}
