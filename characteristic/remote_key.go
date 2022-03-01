package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	RemoteKeyRewind      int = 0
	RemoteKeyFastForward int = 1
	RemoteKeyExit        int = 10
	RemoteKeyPlayPause   int = 11
	RemoteKeyInfo        int = 15
	RemoteKeyNextTrack   int = 2
	RemoteKeyPrevTrack   int = 3
	RemoteKeyArrowUp     int = 4
	RemoteKeyArrowDown   int = 5
	RemoteKeyArrowLeft   int = 6
	RemoteKeyArrowRight  int = 7
	RemoteKeySelect      int = 8
	RemoteKeyBack        int = 9
)

const TypeRemoteKey = "E1"

type RemoteKey struct {
	*Int
}

func NewRemoteKey() *RemoteKey {
	c := NewInt(TypeRemoteKey)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionWrite}
	c.SetMinValue(0)
	c.SetMaxValue(16)
	c.SetStepValue(1)

	return &RemoteKey{c}
}
