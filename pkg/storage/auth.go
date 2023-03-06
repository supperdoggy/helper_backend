package storage

import (
	"context"
	"errors"
	"time"

	"github.com/supperdoggy/helper/pkg/models/dbmodels"
	"github.com/supperdoggy/helper/pkg/utils"
	"go.uber.org/zap"
)

const ValidFor = 30

func (d *mongoClient) NewToken(ctx context.Context, userID string) (string, error) {
	if userID == "" {
		return "", errors.New("id cannot be empty")
	}
	token := utils.GenerateToken()
	d.cache.mut.Lock()
	d.cache.m[token] = dbmodels.Token{
		UserID:   userID,
		TokenStr: token,
		// valid for 30 days
		Expire: time.Now().Add(ValidFor * 24 * time.Hour),
	}
	d.cache.mut.Unlock()
	return token, nil
}

func (d *mongoClient) CheckToken(ctx context.Context, token string) (bool, string) {
	d.cache.mut.Lock()
	v, ok := d.cache.m[token]
	d.cache.mut.Unlock()
	if !ok {
		return false, ""
	}

	d.logger.Info("asdsadsa", zap.Any("tokens", d.cache.m), zap.Any("token", token))

	if time.Now().After(v.Expire) {
		d.cache.mut.Lock()
		delete(d.cache.m, token)
		d.logger.Info("deleting")
		d.cache.mut.Unlock()

		return false, ""
	}
	return true, v.UserID
}
