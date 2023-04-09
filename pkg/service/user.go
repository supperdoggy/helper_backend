package service

import (
	"context"
	"errors"

	"github.com/supperdoggy/helper/pkg/models/dbmodels"
	"github.com/supperdoggy/helper/pkg/storage"
	"github.com/supperdoggy/helper/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	// user
	CreateUser(ctx context.Context, password, email, fullName string) (*dbmodels.User, error)
	DeleteUser(ctx context.Context, id string) (*string, error)
	UpdateUser(ctx context.Context, id, password, email string) (*dbmodels.User, error)

	// auth
	NewToken(ctx context.Context, userID string) (token string, err error)
	CheckToken(ctx context.Context, token string) (userID string, err error)
	Login(ctx context.Context, email, password string) (userID, token string, err error)
	Register(ctx context.Context, email, fullName, password string) (userID, token string, err error)

	// adverts
	CreateAdvert(ctx context.Context, name, body, atype, category, location, userID string, attachments [][]byte) (*dbmodels.Advert, error)
	DeleteAdvert(ctx context.Context, id string) error
	GetAdvert(ctx context.Context, id string) (*dbmodels.Advert, error)
}

type service struct {
	logger *zap.Logger
	db     storage.IMongoClient
}

var (
	ErrBadValues = errors.New("bad values")
)

func NewService(l *zap.Logger, d storage.IMongoClient) IService {
	return &service{
		logger: l,
		db:     d,
	}
}

func (s *service) CreateUser(ctx context.Context, password, email, fullname string) (*dbmodels.User, error) {
	err := utils.ValidateUserEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("error hashing password", zap.Error(err))
		return nil, err
	}

	resp, err := s.db.CreateUser(ctx, email, hashed)
	if err != nil {
		s.logger.Error("error CreateUser", zap.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) (*string, error) {
	err := s.db.DeleteUser(ctx, id)
	if err != nil {
		s.logger.Error("error DeleteUser", zap.Error(err))
		return nil, err
	}

	return &id, nil
}

func (s *service) UpdateUser(ctx context.Context, id, password, email string) (*dbmodels.User, error) {
	err := utils.ValidateUserEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("error hashing password", zap.Error(err))
		return nil, err
	}

	err = s.db.UpdateUser(ctx, id, email, hashed)
	if err != nil {
		s.logger.Error("error updating user", zap.Any("id", id), zap.Error(err))
		return nil, err
	}

	user, err := s.db.GetUser(ctx, id)
	if err != nil {
		s.logger.Error("error getting user", zap.Any("id", id), zap.Error(err))
		return nil, err
	}

	return user, nil
}
