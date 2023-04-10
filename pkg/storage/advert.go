package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/supperdoggy/helper/pkg/models"
	"github.com/supperdoggy/helper/pkg/models/dbmodels"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func (c *mongoClient) CreateAdvert(ctx context.Context, name, body, atype, category, location, userID string, attachments [][]byte) (*dbmodels.Advert, error) {
	u := dbmodels.Advert{
		ID:        uuid.New().String(),
		Name:      name,
		Body:      body,
		Type:      atype,
		Category:  category,
		Location:  location,
		UserID:    userID,
		CreatedAt: time.Now().Unix(),
	}

	_, err := c.advertsCol.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}

	var attachmentsIDs []string
	if len(attachments) != 0 {
		for k, v := range attachments {
			attachmentName := fmt.Sprintf("%s_%v", u.ID, k)
			attachment, err := c.CreateAttachment(ctx, attachmentName, v)
			if err != nil {
				return nil, err
			}
			attachmentsIDs = append(attachmentsIDs, attachment.ID)
		}
	}

	err = c.LinkAttachments(ctx, u.ID, attachmentsIDs)

	return &u, err
}

func (c *mongoClient) LinkAttachments(ctx context.Context, advertID string, attachIDs []string) error {
	_, err := c.advertsCol.UpdateByID(ctx, advertID, bson.M{"$set": bson.M{"attachments": attachIDs}})
	return err
}

func (c *mongoClient) DeleteAdvert(ctx context.Context, id string) error {
	_, err := c.advertsCol.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (c *mongoClient) GetAdvert(ctx context.Context, id string) (*dbmodels.Advert, error) {
	resp := c.advertsCol.FindOne(ctx, bson.M{"_id": id})
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	var user dbmodels.Advert
	if err := resp.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *mongoClient) GetAdverts(ctx context.Context, filter models.AdvertsFilter, limit, offset int) ([]*dbmodels.Advert, error) {
	var adverts []*dbmodels.Advert
	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(offset))

	query := bson.M{}

	if filter.Type != nil {
		query["type"] = *filter.Type
	}
	if filter.Category != nil {
		query["category"] = *filter.Category
	}
	if filter.Location != nil {
		query["location"] = *filter.Location
	}
	if filter.UserID != nil {
		query["user_id"] = *filter.UserID
	}
	if filter.Name != nil {
		query["name"] = bson.M{"$regex": *filter.Name}
	}
	fmt.Println(query)

	resp, err := c.advertsCol.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	if err := resp.All(ctx, &adverts); err != nil {
		return nil, err
	}

	return adverts, nil
}
