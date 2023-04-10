package service

import (
	"context"
	"errors"
	"strings"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) NewToken(ctx context.Context, userID string) (token string, err error) {
	if userID == "" {
		return "", errors.New("no userId provided")
	}

	token, err = s.db.NewToken(ctx, userID)
	if err != nil {
		s.logger.Error("NewToken error creating new token", zap.Error(err))
		return
	}

	return
}

func (s *service) CheckToken(ctx context.Context, token string) (userID string, err error) {
	if token == "" {
		return "", errors.New("Token cannot be empty")
	}

	ok, userID := s.db.CheckToken(ctx, token)
	if !ok {
		return "", errors.New("invalid token")
	}

	return userID, nil
}

func (s *service) Register(ctx context.Context, email, fullName, password string) (userID, token string, err error) {
	email = strings.ToLower(email)

	user, err := s.CreateUser(ctx, password, email, fullName)
	if err != nil {
		s.logger.Error("error Register", zap.Error(err))
		return
	}

	token, err = s.db.NewToken(ctx, user.ID)
	if err != nil {
		s.logger.Error("Register error generating new token", zap.Error(err))
		return
	}

	return user.ID, token, nil
}

func (s *service) Login(ctx context.Context, email, password string) (userID, token string, err error) {

	email = strings.ToLower(email)

	user, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		s.logger.Error("Login error", zap.Error(err))
		return
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		s.logger.Error("error CompareHashAndPassword", zap.Error(err))
		return
	}

	token, err = s.db.NewToken(ctx, user.ID)
	if err != nil {
		s.logger.Error("NewToken error creating new token", zap.Error(err))
		return
	}

	return user.ID, token, nil
}

func (s *service) NewEmailCode(ctx context.Context, email string) error {
	email = strings.ToLower(email)

	code, err := s.emailClient.SendEmail(ctx, email)
	if err != nil {
		s.logger.Error("SendEmail error", zap.Error(err))
		return err
	}

	err = s.db.NewEmailCode(ctx, email, code)
	if err != nil {
		s.logger.Error("NewEmailCode error", zap.Error(err))
		return err
	}

	return nil
}

func (s *service) CheckEmailCode(ctx context.Context, email, code string) error {
	email = strings.ToLower(email)

	ok, err := s.db.CheckEmailCode(ctx, email, code)
	if err != nil {
		s.logger.Error("CheckEmailCode error", zap.Error(err))
		return err
	}

	if !ok {
		return errors.New("invalid code")
	}

	return nil
}
