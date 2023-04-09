package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/supperdoggy/helper/pkg/models/dbmodels"
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
