package mongodb

import (
	"MTSHackathonBackEnd/internal/config"
	"MTSHackathonBackEnd/internal/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoRepository) GetAllCategories() ([]entity.Category, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(
		m.Config.Collections[config.CategoryCollection])

	cur, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cur.Close(context.Background())
	}()

	var categories []entity.Category

	err = cur.All(context.Background(), &categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
