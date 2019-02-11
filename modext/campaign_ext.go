package modext

import (
	"github.com/businessinstincts/traxone/domain"
	"github.com/businessinstincts/traxone/models"
	"github.com/shopspring/decimal"
	"gopkg.in/volatiletech/null.v6"
)

func ConvertCampaignToDomain(campaign *models.Campaign) *domain.Campaign {
	var onspotInfo *domain.OnspotInfo
	if campaign.R == nil || campaign.R.OnspotInfo == nil {
		onspotInfo = nil
	} else {
		onspotInfo = ConvertOnspotInfoToDomain(*campaign.R.OnspotInfo)
	}
	return &domain.Campaign{
		ID:                 campaign.ID,
		Name:               campaign.Name,
		OnspotInfoID:       campaign.OnspotInfoID,
		OnspotInfo:         onspotInfo,
		OrgID:              campaign.OrgID,
		StartDate:          campaign.StartDate,
		EndDate:            campaign.EndDate,
		Type:               domain.CampaignType(campaign.Type),
		Status:             domain.CampaignStatus(campaign.Status),
		PaymentAmount:      nullDecimal(campaign.PaymentAmount),
		PaymentID:          campaign.PaymentID,
		AppNexusCampaignID: campaign.AppnexusCampaignID,
		DeviceCount:        campaign.DeviceCount,
		CreatedAt:          campaign.CreatedAt,
		UpdatedAt:          campaign.UpdatedAt,
	}
}
func nullDecimal(paymentAmount null.Float64) decimal.NullDecimal {
	return decimal.NullDecimal{
		Decimal: decimal.NewFromFloat(paymentAmount.Float64),
		Valid:   paymentAmount.Valid,
	}
}

func ConvertCampaignsToDomain(campaigns *models.CampaignSlice) (*domain.CampaignSlice, error) {
	var cs = domain.CampaignSlice{}
	for _, campaign := range *campaigns {
		c := ConvertCampaignToDomain(campaign)
		cs = append(cs, *c)

	}
	return &cs, nil
}

func ConvertCampaignToModel(campaign domain.Campaign) *models.Campaign {
	return &models.Campaign{
		ID:                 campaign.ID,
		Name:               campaign.Name,
		OnspotInfoID:       campaign.OnspotInfoID,
		AppnexusCampaignID: campaign.AppNexusCampaignID,
		PaymentID:          campaign.PaymentID,
		Type:               int(campaign.Type),
		Status:             int(campaign.Status),
		OrgID:              campaign.OrgID,
		StartDate:          campaign.StartDate,
		EndDate:            campaign.EndDate,
		PaymentAmount:      nullFloat64(campaign.PaymentAmount),
		DeviceCount:        campaign.DeviceCount,
		CreatedAt:          campaign.CreatedAt,
		UpdatedAt:          campaign.UpdatedAt,
	}
}

func nullFloat64(value decimal.NullDecimal) null.Float64 {
	f, _ := value.Decimal.Float64()
	return null.Float64{
		Valid:   value.Valid,
		Float64: f,
	}
}

func ConvertOnspotInfoToModel(info domain.OnspotInfo) *models.OnspotInfo {
	return &models.OnspotInfo{
		ID:          info.ID,
		Information: info.Information,
		CreatedAt:   info.CreatedAt,
		UpdatedAt:   info.UpdatedAt,
	}
}

func ConvertOnspotInfoToDomain(info models.OnspotInfo) *domain.OnspotInfo {
	return &domain.OnspotInfo{
		ID:          info.ID,
		Information: info.Information,
		CreatedAt:   info.CreatedAt,
		UpdatedAt:   info.UpdatedAt,
	}
}
