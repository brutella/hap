package characteristic

import (
	"encoding/base64"
	"reflect"
	"testing"
)

func TestBytesEncoding(t *testing.T) {
	val := []byte{0xFA, 0xAA}
	b := NewBytes(TypeLogs)
	b.Permissions = []string{PermissionWrite, PermissionRead}
	b.SetValue(val)

	expect := base64.StdEncoding.EncodeToString(val)
	if x := b.Val; !reflect.DeepEqual(x, expect) {
		t.Fatal(x)
	}

	if x := b.Value(); !reflect.DeepEqual(x, val) {
		t.Fatal(x)
	}
}
