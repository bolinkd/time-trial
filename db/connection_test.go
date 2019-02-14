package db

import (
	"testing"
)

func TestInterfaces(t *testing.T) {
	var _ DatabaseInterface = (*Connection)(nil)
}
