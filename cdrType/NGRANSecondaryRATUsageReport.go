package cdrType

type NGRANSecondaryRATUsageReport struct { /* Sequence Type */
	NGRANSecondaryRATType *NGRANSecondaryRATType `ber:"tagNum:0,optional"`
	/* Sequence of = 35, FULL Name = struct NGRANSecondaryRATUsageReport__qosFlowsUsageReports */
	/* QosFlowsUsageReport */
	QosFlowsUsageReports []QosFlowsUsageReport `ber:"tagNum:1,optional"`
}
