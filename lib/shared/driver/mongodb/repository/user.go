package repository

import (
	"context"
	"shared/driver/mongodb"
	"shared/driver/mongodb/entity"
	"shared/utils/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserCollection = "users"

type UserRepository struct {
	repository.CrudRepository[entity.User]
	Collection *mongo.Collection
}

func NewUserRepository(db *mongodb.MongoInstance) *UserRepository {
	return &UserRepository{
		Collection:     db.GetCollection(UserCollection),
		CrudRepository: repository.NewCrudRepository[entity.User](UserCollection, db),
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return r.GetEntityBy(ctx, bson.M{"email": email})
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return r.UpdateBSON(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
}

func (r *UserRepository) UpdateLastSeen(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"lastSeen": time.Now()}})
	if err != nil {
		return err
	}
	return nil
}
