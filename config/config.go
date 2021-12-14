package config

import (
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TOKEN = os.Getenv("TOKEN")
var MongoUri = os.Getenv("MONGO_URI")
var DBName = os.Getenv("MONGO_DBNAME")
var AutoCallingCollection = os.Getenv("AUTO_CALLING_COLLECTION")
var TimeOut, _ = time.ParseDuration(os.Getenv("TIME_OUT"))
var KafkaBrokerList = os.Getenv("KAFKA_BROKER_LIST")
var KafkaTopic = os.Getenv("KAFKA_TOPIC")
var KafkaGroupConsumer = os.Getenv("KAFKA_GROUP_CONSUMER")

var mongoDBMaxPoolSize, _ = strconv.ParseUint(os.Getenv("MONGODB_MAX_POOL_SIZE"), 10, 64)
var mongoDBMinPoolSize, _ = strconv.ParseUint(os.Getenv("MONGODB_MIN_POOL_SIZE"), 10, 64)

var DB *mongo.Client
var err error

func ConnectionDB() (*mongo.Client, error) {
	options := options.Client().ApplyURI(MongoUri)
	options.SetMaxPoolSize(mongoDBMaxPoolSize)
	options.SetMaxConnIdleTime(TimeOut)
	options.SetHeartbeatInterval(TimeOut)
	options.SetMinPoolSize(mongoDBMinPoolSize)

	DB, err = mongo.NewClient(options)
	if err != nil {
		logrus.Error("Can't connection to mongodb")
		return nil, err
	}

	return DB, nil
}
