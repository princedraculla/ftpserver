package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const (
	InvolvedPartyPresentNothing int = iota /* No components present */
	InvolvedPartyPresentSIPURI
	InvolvedPartyPresentTELURI
	InvolvedPartyPresentURN
	InvolvedPartyPresentISDNE164
	InvolvedPartyPresentExternalId
)

type InvolvedParty struct {
	Present    int                /* Choice Type */
	SIPURI     *asn.GraphicString `ber:"tagNum:0"`
	TELURI     *asn.GraphicString `ber:"tagNum:1"`
	URN        *asn.GraphicString `ber:"tagNum:2"`
	ISDNE164   *asn.GraphicString `ber:"tagNum:3"`
	ExternalId *asn.UTF8String    `ber:"tagNum:4"`
}
