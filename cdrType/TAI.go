package cdrType

type TAI struct { /* Sequence Type */
	PLMNId PLMNId `ber:"tagNum:0"`
	Tac    TAC    `ber:"tagNum:1"`
}
