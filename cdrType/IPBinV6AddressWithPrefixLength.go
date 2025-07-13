package cdrType

type IPBinV6AddressWithPrefixLength struct { /* Sequence Type */
	IPBinV6Address         IPBinV6Address
	PDPAddressPrefixLength *PDPAddressPrefixLength `ber:"optional,default:64"`
}
