package mongodb

import (
	"MTSHackathonBackEnd/internal/config"
	"MTSHackathonBackEnd/internal/entity"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
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
	_ = cur.Close(context.Background())
	return cert, nil
}

func (m *MongoRepository) ChangeOwnerOfCertificate(id, newUserId string) error {
	coll := m.DB.Database(m.Config.DBName).Collection(
		m.Config.Collections[config.CertificateCollection])

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.M{"$set": bson.M{"owner_id": newUserId}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

func (m *MongoRepository) GenerateLink(id string) string {
	return fmt.Sprintf("UNIQUE LINK + %s", id)
}

func (m *MongoRepository) CreateNewCertificate(cert entity.Certificate) error {
	coll := m.DB.Database(m.Config.DBName).Collection(
		m.Config.Collections[config.CertificateCollection])

	_, err := coll.InsertOne(context.TODO(), &cert)
	return err
}

func (m *MongoRepository) GenerateUniqueId() (string, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(
		m.Config.Collections[config.CertificateCollection])
	count, err := coll.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(count+1, 10), nil
}
