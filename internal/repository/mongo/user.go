package mongodb

import (
	"MTSHackathonBackEnd/internal/config"
	"MTSHackathonBackEnd/internal/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoRepository) GetUserIdByPhoneNumber(phoneNumber string) (string, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(
		m.Config.Collections[config.UserCollection])

	filter := bson.M{"phone_number": phoneNumber}
	cur := coll.FindOne(context.TODO(), filter)

	var user entity.User
	err := cur.Decode(&user)
	return user.PhoneNumber, err
}
