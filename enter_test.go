package couress

import (
	"strings"
	"testing"
)

func TestSetGlobals(t *testing.T) {
	SetGlobals()
}

func TestContains(t *testing.T) {
	t.Log(strings.Contains(">,=,<,>=,<=,between,like", "like"))
}
