package services
import (
	"errors"
	"context"
	"restfulapi/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type UserServiceImpl struct{
	usercollection *mongo.Collection
	ctx				context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(id *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{primitive.E{Key: "id", Value: user.Id}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "id", Value: user.Id}, primitive.E{Key: "name", Value: user.Name}, primitive.E{Key: "dob", Value: user.Dob},primitive.E{Key: "address", Value: user.Address},primitive.E{Key: "description", Value: user.Description},primitive.E{Key: "createdAt", Value: user.CreatedAt}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
