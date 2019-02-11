package domain

import (
	"github.com/businessinstincts/traxone/paytrace"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	orgID = 1
)

func TestCampaign_Validate(t *testing.T) {
	var campaign Campaign

	t.Run("missing start date", func(t *testing.T) {
		assert.Equal(t, ErrNoStartDate, campaign.Validate())
	})
	campaign.StartDate = time.Now()
	t.Run("valid", func(t *testing.T) {
		err := campaign.Validate()
		assert.NoError(t, err)
	})
}

func TestCampaignCreation_Validate(t *testing.T) {
	var campaignCreation CampaignCreation
	t.Run("campaign error", func(t *testing.T) {
		assert.Equal(t, ErrNoStartDate, campaignCreation.Validate())
	})
	campaignCreation.Campaign = Campaign{
		OrgID:     orgID,
		StartDate: time.Now(),
		Type:      CampaignLocation,
		PaymentAmount: decimal.NullDecimal{
			Decimal: decimal.NewFromFloat(1),
			Valid:   true,
		},
	}

	t.Run("Invalid Card information", func(t *testing.T) {
		assert.Equal(t, paytrace.ErrCardExpired, campaignCreation.Validate())
	})

	campaignCreation.TransactionInformation.CreditCard = paytrace.CreditCard{
		EncryptedNumber: "12345",
		ExpirationYear:  "2020",
		ExpirationMonth: "12",
	}

	campaignCreation.TransactionInformation.EncryptedCSC = "1234"
	campaignCreation.TransactionInformation.Amount = 0
	campaignCreation.TransactionInformation.BillingAddress = paytrace.BillingAddress{
		Name:          "test",
		State:         "test",
		ZipCode:       "test",
		StreetAddress: "test",
		City:          "test",
	}

	t.Run("feature collection error", func(t *testing.T) {
		assert.Equal(t, ErrNoFeatures, campaignCreation.Validate())
	})

	campaignCreation.FeatureCollection.Features = []*Feature{
		{
			Properties: map[string]interface{}{},
		},
	}

	t.Run("featureCollection - missing start", func(t *testing.T) {
		err := campaignCreation.Validate()
		assert.Error(t, err)

		_, isTraxErr := err.(TraxError)
		assert.True(t, isTraxErr)
		assert.Equal(t, "missing 'start' in feature[0].Properties", err.Error())
	})

	campaignCreation.FeatureCollection.Features[0].Properties["start"] = 1

	t.Run("featureCollection - start wrong type", func(t *testing.T) {
		err := campaignCreation.Validate()
		assert.Error(t, err)

		_, isTraxErr := err.(TraxError)
		assert.True(t, isTraxErr)
		assert.Equal(t, "'start' in feature[0].Properties is wrong type", err.Error())
	})

	campaignCreation.FeatureCollection.Features[0].Properties["start"] = ""

	t.Run("featureCollection - missing end", func(t *testing.T) {
		err := campaignCreation.Validate()
		assert.Error(t, err)

		_, isTraxErr := err.(TraxError)
		assert.True(t, isTraxErr)
		assert.Equal(t, "missing 'end' in feature[0].Properties", err.Error())
	})

	campaignCreation.FeatureCollection.Features[0].Properties["end"] = 1

	t.Run("featureCollection - end wrong type", func(t *testing.T) {
		err := campaignCreation.Validate()
		assert.Error(t, err)

		_, isTraxErr := err.(TraxError)
		assert.True(t, isTraxErr)
		assert.Equal(t, "'end' in feature[0].Properties is wrong type", err.Error())
	})

	campaignCreation.FeatureCollection.Features[0].Properties["end"] = ""

	t.Run("Valid CampaignCreation - FeatureCollection", func(t *testing.T) {
		assert.NoError(t, campaignCreation.Validate())
	})

	campaignCreation.Campaign.Type = CampaignCustom
	t.Run("Missing Custom Campaign Information", func(t *testing.T) {
		assert.Equal(t, ErrInvalidCampaignInfo, campaignCreation.Validate())
	})

	campaignCreation.CustomInformation = "test information"

	t.Run("Amounts not equal", func(t *testing.T) {
		assert.Equal(t, ErrAmountNotEqual, campaignCreation.Validate())
	})

	campaignCreation.TransactionInformation.Amount = 1

	t.Run("Valid Custom Campaign", func(t *testing.T) {
		assert.NoError(t, campaignCreation.Validate())
	})
}

func TestOnspotInfo_Validate(t *testing.T) {
	var onspotInfo OnspotInfo
	t.Run("Null Information", func(t *testing.T) {
		err := onspotInfo.Validate()
		assert.Equal(t, ErrInformationEmpty, err)
	})
	onspotInfo.Information = "Test"
	t.Run("Valid OnspotInfo", func(t *testing.T) {
		err := onspotInfo.Validate()
		assert.NoError(t, err)
	})
}
