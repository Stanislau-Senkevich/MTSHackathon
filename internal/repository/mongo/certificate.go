package mongodb

import (
	"MTSHackathonBackEnd/internal/config"
	"MTSHackathonBackEnd/internal/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoRepository) GetAllBoughtCertificates(userId string) ([]entity.Certificate, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(
		m.Config.Collections[config.CertificateCollection])

	filter := bson.D{{Key: "owner_id", Value: userId}}

	cur, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cur.Close(context.Background())
	}()

	var cert []entity.Certificate
	err = cur.All(context.Background(), &cert)
	if err != nil {
		return nil, err
	}
	return cert, nil
}
