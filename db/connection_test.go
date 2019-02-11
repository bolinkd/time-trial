package db

import (
	"testing"
)

func TestInterfaces(t *testing.T) {
	var _ DatabaseInterface = (*Connection)(nil)
	var _ CampaignDBInterface = (*Connection)(nil)
	var _ UserDBInterface = (*Connection)(nil)
	var _ OrganizationDBInterface = (*Connection)(nil)
}
