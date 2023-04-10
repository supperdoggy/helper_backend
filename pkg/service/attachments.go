package service

import (
	"context"

	"github.com/supperdoggy/helper/pkg/models/dbmodels"
	"go.uber.org/zap"
)

func (s *service) GetAdvertAttachments(ctx context.Context, advertID string) ([]*dbmodels.Attachment, error) {
	advert, err := s.db.GetAdvert(ctx, advertID)
	if err != nil {
		s.logger.Error("failed to get advert", zap.Error(err))
		return nil, err
	}

	if len(advert.Attachments) == 0 {
		return nil, nil
	}

	attachments, err := s.db.GetAdvertAttachments(ctx, advert.Attachments)
	if err != nil {
		s.logger.Error("failed to get advert attachments", zap.Error(err), zap.Any("advert", advert.Attachments))
		return nil, err
	}

	return attachments, nil
}
