package utils

import (
	"github.com/supperdoggy/helper/pkg/models"
	"github.com/supperdoggy/helper/pkg/models/dbmodels"
)

// map db advert to model advert
func MapDBAdvertToModelAdvert(dbAdvert *dbmodels.Advert) *models.Advert {
	return &models.Advert{
		UserID:      dbAdvert.UserID,
		Name:        dbAdvert.Name,
		Body:        dbAdvert.Body,
		Type:        dbAdvert.Type,
		Category:    dbAdvert.Category,
		Location:    dbAdvert.Location,
		Attachments: dbAdvert.Attachments,
		CreatedAt:   dbAdvert.CreatedAt,
		EditedAt:    dbAdvert.EditedAt,
	}
}

// map db advert to model advert
func MapDBAdvertsToModelAdverts(dbAdverts []*dbmodels.Advert) []*models.Advert {
	var adverts []*models.Advert
	for _, dbAdvert := range dbAdverts {
		adverts = append(adverts, MapDBAdvertToModelAdvert(dbAdvert))
	}

	return adverts
}
