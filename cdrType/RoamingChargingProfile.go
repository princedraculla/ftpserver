package cdrType

type RoamingChargingProfile struct { /* Sequence Type */
	/* Sequence of = 35, FULL Name = struct RoamingChargingProfile__roamingTriggers */
	/* RoamingTrigger */
	RoamingTriggers     []RoamingTrigger     `ber:"tagNum:0,optional"`
	PartialRecordMethod *PartialRecordMethod `ber:"tagNum:1,optional"`
}
