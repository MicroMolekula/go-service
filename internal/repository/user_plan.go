package repository

import (
	"context"
	"errors"
	"github.com/MicroMolekula/gpt-service/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserPlanRepository struct {
	collection *mongo.Collection
}

func NewUserPlanRepository(collection *mongo.Collection) *UserPlanRepository {
	return &UserPlanRepository{
		collection: collection,
	}
}

func (r *UserPlanRepository) CreateOrUpdate(ctx context.Context, userPlan dto.UserPlan) error {
	filter := bson.M{"user_id": userPlan.UserId}

	update := bson.M{
		"$set": bson.M{
			"plan": userPlan.Plan,
		},
		"$setOnInsert": bson.M{
			"user_id":    userPlan.UserId,
			"created_at": time.Now(),
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *UserPlanRepository) GetByUserID(ctx context.Context, userID string) (*dto.UserPlan, error) {
	var userPlan dto.UserPlan
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&userPlan)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &userPlan, nil
}

func (r *UserPlanRepository) AddDayPlan(ctx context.Context, userID string, dayPlan dto.Plan) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$push": bson.M{
			"plan": dayPlan,
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *UserPlanRepository) UpdateDayPlan(ctx context.Context, userID string, day string, updatedPlan dto.Plan) error {
	filter := bson.M{
		"user_id":  userID,
		"plan.day": day,
	}

	update := bson.M{
		"$set": bson.M{
			"plan.$": updatedPlan,
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *UserPlanRepository) RemoveDayPlan(ctx context.Context, userID string, day string) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$pull": bson.M{
			"plan": bson.M{"day": day},
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
