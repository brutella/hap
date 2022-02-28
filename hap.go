package hap

// Status codes for json communication.
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

// Error codes for TLV8 communication.
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
	// HTTPContentTypePairingTLV8 is the HTTP content type for tlv8 data
	HTTPContentTypePairingTLV8 = "application/pairing+tlv8"

	// HTTPContentTypeHAPJson is the HTTP content type for json data
	HTTPContentTypeHAPJson = "application/hap+json"
)

const (
	MethodPair          byte = 0x0 // pair
	MethodPairMFi       byte = 0x1 // MFi compliant accessory
	MethodVerifyPair    byte = 0x2 // verify a pairing
	MethodAddPairing    byte = 0x3 // add client through secure connection
	MethodDeletePairing byte = 0x4 // delete pairing through secure connection
	MethodListPairings  byte = 0x5
)

const (
	// PermissionUser is the user permission for a paired controller.
	PermissionUser byte = 0x0
	// PermissionAdmin is the administrator permission for a paired controller.
	PermissionAdmin byte = 0x1
)
