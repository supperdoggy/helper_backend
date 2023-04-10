package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/supperdoggy/helper/pkg/models/dbmodels"
	"gopkg.in/mgo.v2/bson"
)

func (c *mongoClient) CreateAttachment(ctx context.Context, name string, data []byte) (*dbmodels.Attachment, error) {
	u := dbmodels.Attachment{
		ID:   uuid.New().String(),
		Name: name,
		Data: data,
	}
	_, err := c.attachmentsCol.InsertOne(ctx, u)
	return &u, err
}

func (c *mongoClient) DeleteAttachment(ctx context.Context, id string) error {
	_, err := c.attachmentsCol.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (c *mongoClient) GetAttachment(ctx context.Context, id string) (*dbmodels.Attachment, error) {
	resp := c.attachmentsCol.FindOne(ctx, bson.M{"_id": id})
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	var attachment dbmodels.Attachment
	if err := resp.Decode(&attachment); err != nil {
		return nil, err
	}

	return &attachment, nil
}

func (c *mongoClient) GetAdvertAttachments(ctx context.Context, ids []string) ([]*dbmodels.Attachment, error) {
	var attachments []*dbmodels.Attachment

	for _, v := range ids {
		attachment, err := c.GetAttachment(ctx, v)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, attachment)
	}

	return attachments, nil
}
