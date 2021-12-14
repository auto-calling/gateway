package handler

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

	"github.com/auto-calling/gateway/models"
	"github.com/auto-calling/gateway/utils"
	"github.com/gin-gonic/gin"
)

type MakeEventReq struct {
	Host    string `bson:"host" json:"host" binding:"required"`
	Owner   string `bson:"owner" json:"owner" binding:"required"`
	Service string `bson:"service" json:"service" binding:"required"`
	State   string `bson:"state" json:"state" binding:"required"`
	Status  string `bson:"status" json:"status" binding:"required"`
	Msg     string `bson:"msg" json:"msg" binding:"required"`
	Action  string `bson:"action" json:"action" binding:"required"`
	Created string `bson:"created" json:"created" binding:"required"`
}

type KafkaMsg struct {
	Id interface{} `json:"id"`
	MakeEventReq
}

type MakeEventResBody struct {
	IDs     interface{} `json:"data"`
	Success bool        `json:"success"`
}

func MakeEvent(c *gin.Context) {
	var requestBody MakeEventReq

	// Capture request data from client
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logrus.Error("Failure request data from client")
		RespondWithError(c, http.StatusBadRequest, err)
		return
	}
	// Convert time to Mongodb standard time
	created, err := utils.ConvertStringToTime(requestBody.Created)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	// Restructure logging
	logging := models.Logging{
		ID:      primitive.NewObjectID(),
		Host:    requestBody.Host,
		Owner:   requestBody.Owner,
		Service: requestBody.Host,
		State:   requestBody.State,
		Status:  requestBody.Status,
		Msg:     requestBody.Msg,
		Action:  requestBody.Action,
		Created: created,
	}

	// Write to mongodb
	idLogging, err := models.LoggingInsertOne(logging)
	if err != nil {
		logrus.Error("Can't connection to mongodb")
		RespondWithError(c, http.StatusBadRequest, err)
		return
	}
	// Convert struct to json struct
	b, err := json.Marshal(logging)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	// Convert json to byte
	if err := Producer([]byte(b)); err != nil {
		RespondWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, MakeEventResBody{
		IDs:     idLogging,
		Success: true,
	})
}
