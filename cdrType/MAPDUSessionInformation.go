package cdrType

type MAPDUSessionInformation struct { /* Sequence Type */
	MAPDUSessionIndicator *MAPDUSessionIndicator `ber:"tagNum:0,optional"`
	ATSSSCapability       *ATSSSCapability       `ber:"tagNum:1,optional"`
}
