package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type SMInterface struct { /* Sequence Type */
	InterfaceId   *asn.GraphicString `ber:"tagNum:0,optional"`
	InterfaceText *asn.GraphicString `ber:"tagNum:1,optional"`
	InterfacePort *asn.GraphicString `ber:"tagNum:2,optional"`
	InterfaceType *SMInterfaceType   `ber:"tagNum:3,optional"`
}
