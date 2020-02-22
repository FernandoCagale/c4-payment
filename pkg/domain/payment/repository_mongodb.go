package payment

import (
	"github.com/FernandoCagale/c4-payment/internal/errors"
	"github.com/FernandoCagale/c4-payment/pkg/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION = "payment"
	DATABASE   = "c4-payment-database"
)

type MongodbRepository struct {
	session *mgo.Session
}

func NewMongodbRepository(session *mgo.Session) *MongodbRepository {
	return &MongodbRepository{session}
}

func (repo *MongodbRepository) FindAll() (payments []*entity.Customer, err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	err = coll.Find(nil).All(&payments)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return payments, nil
}

func (repo *MongodbRepository) FindById(ID string) (payment *entity.Customer, err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	err = coll.FindId(ID).One(&payment)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			return nil, errors.ErrNotFound
		default:
			return nil, errors.ErrInternalServer
		}
	}

	return payment, nil
}

func (repo *MongodbRepository) Create(e *entity.Customer) (err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	change, err := repo.FindById(e.Code)
	switch err {
	case errors.ErrNotFound:
		err = coll.Insert(e)
	case nil:
		err = coll.Update(change, bson.M{"$push": bson.M{"payments": bson.M{"$each": e.Payments}}})
	default:
		return errors.ErrInternalServer
	}

	if err != nil {
		if mgo.IsDup(err) {
			return errors.ErrConflict
		}
		return errors.ErrInternalServer
	}
	return nil
}

func (repo *MongodbRepository) DeleteById(ID string) (err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	err = coll.RemoveId(ID)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			return errors.ErrNotFound
		default:
			return errors.ErrInternalServer
		}
	}

	return nil
}
