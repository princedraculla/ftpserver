package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type NodeAddress struct {
	Present    int                /* Choice Type */
	IPAddress  *IPAddress         `ber:"tagNum:0"`
	DomainName *asn.GraphicString `ber:"tagNum:1"`
}
