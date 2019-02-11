package domain

import (
	"errors"
	"fmt"
	"github.com/businessinstincts/traxone/paytrace"
	"github.com/shopspring/decimal"
	"gopkg.in/volatiletech/null.v6"
	"time"
)

var (
	ErrInformationEmpty    = errors.New("information can't be null")
	ErrNoOrganization      = errors.New("organization ID is missing")
	ErrNoStartDate         = errors.New("start date is missing")
	ErrAmountNotEqual      = errors.New("transaction and campaign amount must be equal")
	ErrInvalidCampaignInfo = errors.New("campaign information cannot be empty")
)

const (
	CampaignLocation CampaignType = iota
	CampaignBusiness
	CampaignCustom
)

const (
	CampaignNotPaid CampaignStatus = iota
	CampaignPaid
	CampaignInProgress
	CampaignCompleted
	CampaignCancelled
)

type CampaignType int
type CampaignStatus int
type CampaignSlice []Campaign

type OnspotInfo struct {
	ID          int       `json:"id"`
	Information string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CampaignCreation struct {
	Campaign               Campaign          `json:"campaign"`
	FeatureCollection      FeatureCollection `json:"feature_collection"`
	CustomInformation      string            `json:"custom_info"`
	TransactionInformation paytrace.Request  `json:"transaction_info"`
}

func (campaignCreation CampaignCreation) Validate() error {
	err := campaignCreation.Campaign.Validate()
	if err != nil {
		return err
	}
	if campaignCreation.Campaign.Type == CampaignLocation || campaignCreation.Campaign.Type == CampaignBusiness {
		err = campaignCreation.TransactionInformation.Validate()
		if err != nil {
			return err
		}
		return campaignCreation.ValidateFeatureCollection()

	} else {
		if len(campaignCreation.CustomInformation) == 0 {
			return ErrInvalidCampaignInfo
		}
	}
	amount1, _ := campaignCreation.Campaign.PaymentAmount.Decimal.Float64()

	amount2 := campaignCreation.TransactionInformation.Amount
	if amount1-amount2 > 0.1 {
		return ErrAmountNotEqual
	}
	return nil
}

func (campaignCreation CampaignCreation) ValidateFeatureCollection() error {
	collection := campaignCreation.FeatureCollection
	err := collection.Validate()
	if err != nil {
		return err
	}
	for i, feature := range collection.Features {
		errMissingString := "missing '%s' in feature[%d].Properties"
		errInvalidType := "'%s' in feature[%d].Properties is wrong type"
		if _, ok := feature.Properties["start"]; !ok {
			err = errors.New(fmt.Sprintf(errMissingString, "start", i))
		} else if _, ok = feature.Properties["start"].(string); !ok {
			err = errors.New(fmt.Sprintf(errInvalidType, "start", i))
		} else if _, ok := feature.Properties["end"]; !ok {
			err = errors.New(fmt.Sprintf(errMissingString, "end", i))
		} else if _, ok = feature.Properties["end"].(string); !ok {
			err = errors.New(fmt.Sprintf(errInvalidType, "end", i))
		}
	}
	return err
}

type Campaign struct {
	ID                 int                 `json:"id"`
	Name               null.String         `json:"name"`
	StartDate          time.Time           `json:"start_date"`
	EndDate            null.Time           `json:"end_date"`
	PaymentID          null.String         `json:"payment_id"`
	PaymentAmount      decimal.NullDecimal `json:"payment_amount"`
	OnspotInfoID       null.Int            `json:"onspot_info_id"`
	OnspotInfo         *OnspotInfo         `json:"onspot_info,omitempty"`
	OrgID              int                 `json:"org_id"`
	Type               CampaignType        `json:"type"`
	AppNexusCampaignID null.String         `json:"appnexus_campaign_id"`
	Status             CampaignStatus      `json:"status"`
	DeviceCount        null.Int            `json:"device_count"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
}

func (c Campaign) Validate() error {
	if c.StartDate.IsZero() {
		return ErrNoStartDate
	}
	return nil
}

func (i OnspotInfo) Validate() error {
	if i.Information == "" {
		return ErrInformationEmpty
	}
	return nil
}
