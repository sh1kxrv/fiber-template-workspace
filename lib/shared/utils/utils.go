package utils

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateContextTimeout(timeoutInSec int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(timeoutInSec)*time.Second)
}

func PartialUpdateDocument[T any](ctx context.Context, collection *mongo.Collection, documentID primitive.ObjectID, updates interface{}) error {
	updatesValue := reflect.ValueOf(updates)
	if updatesValue.Kind() != reflect.Map && updatesValue.Kind() != reflect.Struct {
		return fmt.Errorf("updates must be a struct or map")
	}

	updateData, err := bson.Marshal(updates)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": documentID}
	update := bson.M{"$set": updateData}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func GetIDFromObject[T any](obj T, fieldName string) (string, error) {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return "", errors.New("obj must be a struct")
	}

	field := val.FieldByName(fieldName)
	if !field.IsValid() || field.Kind() != reflect.String {
		return "", fmt.Errorf("field '%s' is not a string", fieldName)
	}

	return field.String(), nil
}

func ParseIDsFromObject[T any](obj []T, fieldName string) ([]primitive.ObjectID, error) {
	var parsedIds []primitive.ObjectID
	for _, o := range obj {
		id, err := GetIDFromObject[T](o, fieldName)
		if err != nil {
			return nil, err
		}
		parsedId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		parsedIds = append(parsedIds, parsedId)
	}
	return parsedIds, nil
}

func ParseIDsFromString(ids []string) ([]primitive.ObjectID, error) {
	var parsedIds []primitive.ObjectID
	for _, id := range ids {
		parsedId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		parsedIds = append(parsedIds, parsedId)
	}
	return parsedIds, nil
}

func Find[T any](arr []T, predicate func(T) bool) (T, error) {
	var zero T
	for _, v := range arr {
		if predicate(v) {
			return v, nil
		}
	}
	return zero, errors.New("no matching element found")
}
