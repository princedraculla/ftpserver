package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type GNbId struct { /* Sequence Type */
	BitLength int64         `ber:"tagNum:0"`
	GNbValue  asn.IA5String `ber:"tagNum:1"`
}
