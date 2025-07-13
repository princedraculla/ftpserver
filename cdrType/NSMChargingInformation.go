package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type NSMChargingInformation struct { /* Set Type */
	ManagementOperation    *ManagementOperation `ber:"tagNum:0,optional"`
	IDnetworkSliceInstance *asn.OctetString     `ber:"tagNum:1,optional"`
	/* Sequence of = 35, FULL Name = struct NSMChargingInformation__listOfserviceProfileChargingInformation */
	/* ServiceProfileChargingInformation */
	ListOfserviceProfileChargingInformation []ServiceProfileChargingInformation `ber:"tagNum:2,optional"`
	ManagementOperationStatus               *ManagementOperationStatus          `ber:"tagNum:3,optional"`
	OperationalState                        *OperationalState                   `ber:"tagNum:4,optional"`
	AdministrativeState                     *AdministrativeState                `ber:"tagNum:5,optional"`
}
