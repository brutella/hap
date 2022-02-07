package hap

const (
	JsonStatusSuccess                     = 0
	JsonStatusInsufficientPrivileges      = -70401
	JsonStatusServiceCommunicationFailure = -70402
	JsonStatusResourceBusy                = -70403
	JsonStatusReadOnlyCharacteristic      = -70404
	JsonStatusWriteOnlyCharacteristic     = -70405
	JsonStatusNotificationNotSupported    = -70406
	JsonStatusOutOfResource               = -70407
	JsonStatusOperationTimedOut           = -70408
	JsonStatusResourceDoesNotExist        = -70409
	JsonStatusInvalidValueInRequest       = -70410
)

const (
	TlvErrorUnknown        = 0x1
	TlvErrorInvalidRequest = 0x2
	TlvErrorAuthentication = 0x2
	TlvErrorBackoff        = 0x3
	TlvErrorMaxPeers       = 0x4
	TlvErrorUnknownPeer    = 0x4
	TlvErrorMaxTries       = 0x5
	TlvErrorUnavailable    = 0x6
	TlvErrorBusy           = 0x7
)

const (
	// HTTPContentTypePairingTLV8 is the HTTP content type for pairing
	HTTPContentTypePairingTLV8 = "application/pairing+tlv8"

	// HTTPContentTypeHAPJson is the HTTP content type for json data
	HTTPContentTypeHAPJson = "application/hap+json"
)
