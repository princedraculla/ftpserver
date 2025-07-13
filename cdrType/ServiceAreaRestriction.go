package cdrType

type ServiceAreaRestriction struct { /* Sequence Type */
	RestrictionType *RestrictionType `ber:"tagNum:0,optional"`
	/* Sequence of = 35, FULL Name = struct ServiceAreaRestriction__areas */
	/* Area */
	Areas                         []Area `ber:"tagNum:1,optional"`
	MaxNumOfTAs                   *int64 `ber:"tagNum:2,optional"`
	MaxNumOfTAsForNotAllowedAreas *int64 `ber:"tagNum:3,optional"`
}
