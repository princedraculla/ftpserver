package cdrType

type UserLocationInformationStructured struct { /* Sequence Type */
	EutraLocation *EutraLocation `ber:"tagNum:0,optional"`
	NrLocation    *NrLocation    `ber:"tagNum:1,optional"`
	N3gaLocation  *N3gaLocation  `ber:"tagNum:2,optional"`
}
