package links

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"gitlab.com/robotomize/gb-golang/homework/03-01-umanager/internal/database"
)

const collection = "links"

func New(db *mongo.Database, timeout time.Duration) *Repository {
	return &Repository{db: db, timeout: timeout}
}

type Repository struct {
	db      *mongo.Database
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateReq) (database.Link, error) {
	l := database.Link{
		ID:     req.ID,
		URL:    req.URL,
		Title:  req.Title,
		Tags:   req.Tags,
		Images: req.Images,
		UserID: req.UserID,
	}

	if _, err := r.db.Collection(collection).InsertOne(ctx, l); err != nil {
		return database.Link{}, err
	}

	return l, nil
}

func (r *Repository) FindByUserAndURL(ctx context.Context, link, userID string) (database.Link, error) {
	var l database.Link

	filter := bson.M{"url": link, "userID": userID}
	if err := r.db.Collection(collection).FindOne(ctx, filter).Decode(&l); err != nil {
		return database.Link{}, err
	}

	return l, nil
}

func (r *Repository) FindByCriteria(ctx context.Context, criteria Criteria) ([]database.Link, error) {
	var links []database.Link

	// Подготовка фильтра
	filter := bson.M{}
	if criteria.UserID != nil {
		filter["userId"] = *criteria.UserID
	}
	if len(criteria.Tags) > 0 {
		filter["tags"] = bson.M{"$all": criteria.Tags}
	}

	// Подготовка опций запроса
	findOptions := options.Find()
	if criteria.Limit != nil {
		findOptions.SetLimit(*criteria.Limit)
	}
	if criteria.Offset != nil {
		findOptions.SetSkip(*criteria.Offset)
	}

	// Выполнение запроса
	cursor, err := r.db.Collection(collection).Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var l database.Link
		if err := cursor.Decode(&l); err != nil {
			return nil, err
		}
		links = append(links, l)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return links, nil
}
