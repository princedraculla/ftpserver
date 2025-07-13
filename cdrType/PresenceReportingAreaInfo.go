package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type PresenceReportingAreaInfo struct { /* Sequence Type */
	PresenceReportingAreaIdentifier   asn.OctetString                    `ber:"tagNum:0"`
	PresenceReportingAreaStatus       *PresenceReportingAreaStatus       `ber:"tagNum:1,optional"`
	PresenceReportingAreaElementsList *PresenceReportingAreaElementsList `ber:"tagNum:2,optional"`
	PresenceReportingAreaNode         *PresenceReportingAreaNode         `ber:"tagNum:3,optional"`
}
