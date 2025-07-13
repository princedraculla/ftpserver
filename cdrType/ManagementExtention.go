package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type ManagementExtension struct { /* Sequence Type */
	Identifier   asn.ObjectIdentifier
	Significance *bool                          `ber:"tagNum:1,optional,default:FALSE"`
	Information  ManagementExtensionInformation `ber:"tagNum:2"`
}

const (
	ManagementExtensionPresentNothing int = iota /* No components present */
)

type ManagementExtensionInformation struct {
	Present int /* Open Type */
}
