package storage

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/supperdoggy/helper/pkg/models/dbmodels"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

type obj map[string]interface{}

type tokenCache struct {
	m   map[string]dbmodels.Token
	mut sync.Mutex
}

type IMongoClient interface {
	// user
	CreateUser(ctx context.Context, email string, password []byte) (*dbmodels.User, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id, email string, password []byte) error
	GetUser(ctx context.Context, id string) (*dbmodels.User, error)
	GetUserByEmail(ctx context.Context, email string) (*dbmodels.User, error)

	// auth
	NewToken(ctx context.Context, userID string) (string, error)
	CheckToken(ctx context.Context, token string) (bool, string)
}

type mongoClient struct {
	logger   *zap.Logger
	client   *mongo.Client
	usersCol *mongo.Collection

	cache tokenCache
}

func NewMongoClient(ctx context.Context, url string, l *zap.Logger) (IMongoClient, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		l.Error("error connecting to mongodb", zap.Error(err))
		return nil, err
	}

	return &mongoClient{
		client: client,
		logger: l,
		cache: tokenCache{
			m: make(map[string]dbmodels.Token),
		},

		usersCol: client.Database("helper").Collection("users"),
	}, nil
}

func (c *mongoClient) CreateUser(ctx context.Context, email string, password []byte) (*dbmodels.User, error) {
	u := dbmodels.User{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Unix(),
	}
	_, err := c.usersCol.InsertOne(ctx, u)
	return &u, err
}

func (c *mongoClient) DeleteUser(ctx context.Context, id string) error {
	_, err := c.usersCol.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (c *mongoClient) UpdateUser(ctx context.Context, id, email string, password []byte) error {
	_, err := c.usersCol.UpdateByID(ctx, id, bson.M{"$set": bson.M{"email": email, "password": password, "edited_at": time.Now().Unix()}})
	return err
}

func (c *mongoClient) GetUser(ctx context.Context, id string) (*dbmodels.User, error) {
	resp := c.usersCol.FindOne(ctx, bson.M{"_id": id})
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	var user dbmodels.User
	if err := resp.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil

}

func (c *mongoClient) GetUserByEmail(ctx context.Context, email string) (*dbmodels.User, error) {
	resp := c.usersCol.FindOne(ctx, bson.M{"email": email})
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	var user dbmodels.User
	if err := resp.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
