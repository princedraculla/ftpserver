package cdrType

import "github.com/princedraculla/ftpservertest/asn"

type SubscriptionID struct { /* Set Type */
	SubscriptionIDType SubscriptionIDType `ber:"tagNum:0"`
	SubscriptionIDData asn.UTF8String     `ber:"tagNum:1"`
}
