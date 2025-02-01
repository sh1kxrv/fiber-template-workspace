package repository

import (
	"common/driver/mongodb"
	"common/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CrudRepository[T any] struct {
	Collection *mongo.Collection
}

func NewCrudRepository[T any](name string, db *mongodb.MongoInstance) CrudRepository[T] {
	return CrudRepository[T]{
		Collection: db.GetCollection(name),
	}
}

func (r *CrudRepository[T]) CreateEntity(ctx context.Context, entity *T) (*T, error) {
	_, err := r.Collection.InsertOne(ctx, entity)
	return entity, err
}

func (r *CrudRepository[T]) GetEntityByID(ctx context.Context, id primitive.ObjectID) (*T, error) {
	var entity T
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&entity)
	return &entity, err
}

func (r *CrudRepository[T]) UpdateBSON(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *CrudRepository[T]) GetAll(ctx context.Context, filter bson.M, limit, offset int64) ([]T, error) {
	return utils.CursoredFind[T](r.Collection, ctx, filter, limit, offset)
}

func (r *CrudRepository[T]) GetEntityBy(ctx context.Context, filter interface{}) (*T, error) {
	var entity T
	err := r.Collection.FindOne(ctx, filter).Decode(&entity)
	return &entity, err
}
