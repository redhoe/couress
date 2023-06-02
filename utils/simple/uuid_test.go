package simple

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestUuid(t *testing.T) {
	for i := 0; i < 100; i++ {
		uuid := uuid.NewV4()
		t.Log(uuid.Bytes())
		t.Log(uuid.String())
		t.Log(uuid)
	}
}

func TestString2Uuid(t *testing.T) {
	uuid := StringFormatUuid("123e1567-e89b-12d3-a456-426655440000")
	t.Log(uuid.Bytes())
}
