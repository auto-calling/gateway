package models

import (
	"context"
	"github.com/auto-calling/gateway/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Logging struct {
	ID      primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Host    string             `bson:"host" json:"host" binding:"required"`
	Owner   string             `bson:"owner" json:"owner" binding:"required"`
	Service string             `bson:"service" json:"service" binding:"required"`
	State   string             `bson:"state" json:"state" binding:"required"`
	Status  string             `bson:"status" json:"status" binding:"required"`
	Msg     string             `bson:"msg" json:"msg" binding:"required"`
	Action  string             `bson:"action" json:"action" binding:"required"`
	Created time.Time          `bson:"created" json:"created" binding:"required"`
}

func LoggingInsertOne(callLogs interface{}) (interface{}, error) {
	collection := config.DB.Database(config.DBName).Collection(config.AutoCallingCollection)

	insertResult, err := collection.InsertOne(context.TODO(), callLogs)
	if err != nil {
		logrus.Error("Failure import logging to mongodb")
		return nil, err
	}

	return insertResult.InsertedID, err
}
