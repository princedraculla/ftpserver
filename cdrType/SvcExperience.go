package cdrType

type SvcExperience struct { /* Sequence Type */
	Mos        *int64 `ber:"tagNum:0,optional"`
	UpperRange *int64 `ber:"tagNum:1,optional"`
	LowerRange *int64 `ber:"tagNum:2,optional"`
}
