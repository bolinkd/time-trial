package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidGroupClubID = errors.New("invalid group")
)

type Group struct {
	*models.Group
}

type GroupSlice struct {
	models.GroupSlice
}

func (g Group) Validate() error {
	if !g.OrganizationID.Valid {
		return ErrInvalidGroupClubID
	}
	return nil
}

func (g *Group) MarshalJSON() ([]byte, error) {
	if g.R == nil {
		return json.Marshal(g.Group)
	}
	return json.Marshal(&struct {
		*models.Group
		Rowers    models.RowerSlice `json:"rowers,omitempty"`
		Parent    *models.Group     `json:"parent,omitempty"`
		SubGroups models.GroupSlice `json:"sub_groups,omitempty"`
	}{
		Group:     g.Group,
		Parent:    g.R.Parent,
		SubGroups: g.R.ParentGroups,
	})
}

func (gs GroupSlice) MarshalJSON() ([]byte, error) {
	gsd := make([]*Group, 0)
	for _, g := range gs.GroupSlice {
		gsd = append(gsd, &Group{g})
	}
	return json.Marshal(gsd)
}
