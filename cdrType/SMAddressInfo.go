package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type SMAddressInfo struct { /* Sequence Type */
	SMAddressType   *SMAddressType     `ber:"tagNum:0,optional"`
	SMAddressData   *asn.GraphicString `ber:"tagNum:1,optional"`
	SMAddressDomain *SMAddressDomain   `ber:"tagNum:2,optional"`
}
