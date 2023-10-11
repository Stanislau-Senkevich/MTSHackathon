package mongodb

import (
	"MTSHackathonBackEnd/internal/config"
	"MTSHackathonBackEnd/internal/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoRepository) GetAllServices() ([]entity.Service, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(
		m.Config.Collections[config.ServiceCollection])

	cur, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cur.Close(context.Background())
	}()

	var services []entity.Service

	err = cur.All(context.Background(), &services)
	if err != nil {
		return nil, err
	}
	return services, nil
}
