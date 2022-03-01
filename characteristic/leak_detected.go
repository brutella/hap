package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	LeakDetectedLeakNotDetected int = 0
	LeakDetectedLeakDetected    int = 1
)

const TypeLeakDetected = "70"

type LeakDetected struct {
	*Int
}

func NewLeakDetected() *LeakDetected {
	c := NewInt(TypeLeakDetected)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &LeakDetected{c}
}
