package cdrType

type Ecgi struct { /* Sequence Type */
	PlmnId      PLMNId      `ber:"tagNum:0"`
	EutraCellId EutraCellId `ber:"tagNum:1"`
	Nid         *Nid        `ber:"tagNum:2,optional"`
}
