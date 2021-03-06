package mongo_pool_service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetTimeRangeFilterWithPatientUID(patientUID string, from string, to string) (bson.M, error) {
	layout := "2006-01-02T15:04:05"
	parsedFrom, parseFromErr := time.Parse(layout, from)
	if parseFromErr != nil {
		return bson.M{}, parseFromErr
	}
	parsedTo, parseToErr := time.Parse(layout, to)
	if parseFromErr != nil {
		return bson.M{}, parseToErr
	}

	return bson.M{
		"patientuid": patientUID,
		"timestamp":  bson.M{"$gte": parsedFrom, "$lte": parsedTo},
	}, nil
}

func GetTimeRangeFilter(from string, to string) (bson.M, error) {
	layout := "2006-01-02T15:04:05"
	parsedFrom, parseFromErr := time.Parse(layout, from)
	if parseFromErr != nil {
		return bson.M{}, parseFromErr
	}
	parsedTo, parseToErr := time.Parse(layout, to)
	if parseFromErr != nil {
		return bson.M{}, parseToErr
	}

	return bson.M{
		"timestamp": bson.M{"$gte": parsedFrom, "$lte": parsedTo},
	}, nil
}

func GetUIDFilter(uid string) (bson.M, error) {
	return bson.M{
		"uid": uid,
	}, nil
}
