package cdrType

type ServingNetworkFunctionID struct { /* Sequence Type */
	ServingNetworkFunctionInformation NetworkFunctionInformation `ber:"tagNum:0"`
	AMFIdentifier                     *AMFID                     `ber:"tagNum:1,optional"`
}
