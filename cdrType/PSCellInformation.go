package cdrType

type PSCellInformation struct { /* Sequence Type */
	NRcgi *Ncgi `ber:"tagNum:0,optional"`
	Ecgi  *Ecgi `ber:"tagNum:1,optional"`
}
