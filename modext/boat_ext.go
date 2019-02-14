package modext

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/null"
)

func ConvertBoatsToDomain(boats models.BoatSlice) domain.BoatSlice {
	var bs = domain.BoatSlice{}
	for _, boat := range boats {
		b := ConvertBoatToDomain(boat)
		bs = append(bs, *b)

	}
	return bs
}

func ConvertBoatToDomain(boat *models.Boat) *domain.Boat {
	return &domain.Boat{
		ID:          boat.ID,
		TimeTrialID: boat.TimeTrialID,
		Name:        null.StringFrom(boat.Name),
		Start:       boat.Start,
		End:         boat.End,
		Time:        boat.Time,
		BowMarker:   boat.BowMarker,
		CreatedAt:   boat.CreatedAt,
		UpdatedAt:   boat.UpdatedAt,
	}
}

func ConvertBoatToModel(boat domain.Boat) *models.Boat {
	return &models.Boat{
		ID:          boat.ID,
		TimeTrialID: boat.TimeTrialID,
		Name:        boat.Name.String,
		Start:       boat.Start,
		End:         boat.End,
		Time:        boat.Time,
		BowMarker:   boat.BowMarker,
		CreatedAt:   boat.CreatedAt,
		UpdatedAt:   boat.UpdatedAt,
	}
}
