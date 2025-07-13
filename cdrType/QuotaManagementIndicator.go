package cdrType

import "github.com/princedraculla/ftpservertest/asn"

const ( /* Enum Type */
	QuotaManagementIndicatorPresentOnlineCharging           asn.Enumerated = 0
	QuotaManagementIndicatorPresentOfflineCharging          asn.Enumerated = 1
	QuotaManagementIndicatorPresentQuotaManagementSuspended asn.Enumerated = 2
)

type QuotaManagementIndicator struct {
	Value asn.Enumerated
}
