package cdrType

type NSSAIMap struct { /* Sequence Type */
	ServingSnssai SingleNSSAI `ber:"tagNum:0"`
	HomeSnssai    SingleNSSAI `ber:"tagNum:1"`
}
