package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type SMAddressDomain struct { /* Sequence Type */
	SMDomainName       *asn.GraphicString `ber:"tagNum:0,optional"`
	ThreeGPPIMSIMCCMNC *PLMNId            `ber:"tagNum:1,optional"`
}
