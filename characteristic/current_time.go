package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCurrentTime = "9B"

type CurrentTime struct {
	*String
}

func NewCurrentTime() *CurrentTime {
	c := NewString(TypeCurrentTime)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead, PermissionWrite}
	c.Val = ""

	return &CurrentTime{c}
}
