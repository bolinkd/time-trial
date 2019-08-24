package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidGroup = errors.New("invalid group")
)

type Group struct {
	*models.Group
}

type GroupSlice struct {
	models.GroupSlice
}

func (g Group) Validate() error {
	if !g.ClubID.Valid {
		return ErrInvalidGroup
	}
	return nil
}

func (g *Group) MarshalJSON() ([]byte, error) {
	if g.R == nil {
		return json.Marshal(g.Group)
	}
	return json.Marshal(&struct {
		*models.Group
		Rowers models.RowerSlice `json:"rowers,omitempty"`
		Club   *models.Club      `json:"club,omitempty"`
	}{
		Group:  g.Group,
		Rowers: g.R.GroupRowers,
		Club:   g.R.Club,
	})
}

func (gs GroupSlice) MarshalJSON() ([]byte, error) {
	var gsd []*Group
	for _, g := range gs.GroupSlice {
		gsd = append(gsd, &Group{g})
	}
	return json.Marshal(gsd)
}
