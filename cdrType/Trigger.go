package cdrType

type Trigger struct {
	Present    int         /* Choice Type */
	SMFTrigger *SMFTrigger `ber:"tagNum:0"`
}
