package modext

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gopkg.in/volatiletech/null.v6"
	"testing"

	"github.com/businessinstincts/traxone/domain"
	"github.com/businessinstincts/traxone/models"
)

func TestConvertCampaignToDomain(t *testing.T) {
	campaignD := &domain.Campaign{}
	campaignM := &models.Campaign{}

	t.Run("", func(t *testing.T) {
		actual := ConvertCampaignToDomain(campaignM)
		assert.Equal(t, campaignD, actual)
	})
}

func Test_nullDecimal(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		d := nullDecimal(null.Float64From(73))
		assert.Equal(t, true, d.Valid)
	})
	t.Run("Null", func(t *testing.T) {
		d := nullDecimal(null.NewFloat64(73, false))
		assert.Equal(t, false, d.Valid)
	})
}

func TestConvertCampaignsToDomain(t *testing.T) {
	campaignsD := &domain.CampaignSlice{
		{},
	}
	campaignsM := &models.CampaignSlice{
		{},
	}

	actual, err := ConvertCampaignsToDomain(campaignsM)

	assert.NoError(t, err)
	assert.Equal(t, campaignsD, actual)
}

func TestConvertCampaignToModel(t *testing.T) {
	campaignD := domain.Campaign{}
	campaignM := &models.Campaign{}

	actual := ConvertCampaignToModel(campaignD)

	assert.Equal(t, campaignM, actual)
}

func Test_nullFloat64(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		f := nullFloat64(decimal.NullDecimal{Decimal: decimal.NewFromFloat(73), Valid: true})
		assert.Equal(t, true, f.Valid)
	})
	t.Run("Null", func(t *testing.T) {
		f := nullFloat64(decimal.NullDecimal{Decimal: decimal.NewFromFloat(73), Valid: false})
		assert.Equal(t, false, f.Valid)
	})
}

func TestConvertOnspotInfoToModel(t *testing.T) {
	onspotInfoD := domain.OnspotInfo{}
	onspotInfoM := &models.OnspotInfo{}

	actual := ConvertOnspotInfoToModel(onspotInfoD)

	assert.Equal(t, onspotInfoM, actual)
}

func TestConvertOnspotInfoToDomain(t *testing.T) {
	onspotInfoD := &domain.OnspotInfo{}
	onspotInfoM := models.OnspotInfo{}

	actual := ConvertOnspotInfoToDomain(onspotInfoM)

	assert.Equal(t, onspotInfoD, actual)
}
