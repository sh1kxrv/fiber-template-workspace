package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CursoredFind[T any](collection *mongo.Collection, context context.Context, filter interface{}, limit, offset int64) ([]T, error) {
	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	cursor, err := collection.Find(context, filter, options)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context)
	return cursorAll[T](cursor, context)
}

func CursoredAggregate[T any](collection *mongo.Collection, context context.Context, pipe mongo.Pipeline) ([]T, error) {
	cursor, err := collection.Aggregate(context, pipe)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context)
	return cursorAll[T](cursor, context)
}

func cursorAll[T any](cursor *mongo.Cursor, ctx context.Context) ([]T, error) {
	var entities []T = make([]T, 0)
	err := cursor.All(ctx, &entities)
	if err != nil {
		return nil, err
	}

	return entities, nil
}
