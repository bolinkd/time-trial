package socket

import "testing"

func TestInterfaces(t *testing.T) {
	var _ ClientInterface = (*Client)(nil)
}
