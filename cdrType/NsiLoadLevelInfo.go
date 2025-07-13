package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type NsiLoadLevelInfo struct { /* Sequence Type */
	LoadLevelInformation *int64           `ber:"tagNum:0,optional"`
	Snssai               *SingleNSSAI     `ber:"tagNum:1,optional"`
	NsiId                *asn.OctetString `ber:"tagNum:2,optional"`
}
