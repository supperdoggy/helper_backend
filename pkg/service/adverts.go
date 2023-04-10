package service

import (
	"context"

	"github.com/supperdoggy/helper/pkg/models"
	"github.com/supperdoggy/helper/pkg/models/dbmodels"
	"go.uber.org/zap"
)

func (s *service) CreateAdvert(ctx context.Context, name, body, atype, category, location, userID string, attachments [][]byte) (*dbmodels.Advert, error) {
	advert, err := s.db.CreateAdvert(ctx, name, body, atype, category, location, userID, attachments)
	if err != nil {
		s.logger.Error("failed to create advert", zap.Error(err))
		return nil, err
	}
	return advert, nil
}

func (s *service) DeleteAdvert(ctx context.Context, id string) error {
	err := s.db.DeleteAdvert(ctx, id)
	if err != nil {
		s.logger.Error("failed to delete advert", zap.Error(err))
		return err
	}
	return nil
}

func (s *service) GetAdvert(ctx context.Context, id string) (*dbmodels.Advert, error) {
	advert, err := s.db.GetAdvert(ctx, id)
	if err != nil {
		s.logger.Error("failed to get adverts", zap.Error(err))
		return nil, err
	}
	return advert, nil
}

func (s *service) GetAdverts(ctx context.Context, filter models.AdvertsFilter, limit, offset int) ([]*dbmodels.Advert, error) {
	adverts, err := s.db.GetAdverts(ctx, filter, limit, offset)
	if err != nil {
		s.logger.Error("failed to get adverts", zap.Error(err))
		return nil, err
	}
	return adverts, nil
}
