package service_logger

import (
	"errors"
	"fmt"

	bson "go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slices"
)

type UnmarshalledMessage struct {
	Database   string
	Collection string
	UUID       string
	Payload    map[string]interface{}
}

const (
	Database   = "database"
	Collection = "collection"
	UUID       = "uuid"
	Payload    = "payload"
)

func unmarshalBaseMessage(message []byte) (UnmarshalledMessage, error) {

	doc := make(map[string]interface{})
	if err := bson.UnmarshalExtJSON(message, true, &doc); err != nil {
		fmt.Printf("Unmarshal failed: %v\n", err)
		return UnmarshalledMessage{}, err
	}

	if err := validateUnmarshalledMessage(&doc); err != nil {
		return UnmarshalledMessage{}, err
	}

	if err := validateAllowedDatabase(doc[Database].(string)); err != nil {
		return UnmarshalledMessage{}, err
	}

	uuid := doc[UUID].(string)
	payload := doc[Payload].(map[string]interface{})

	// fill uuid in payload
	payload[UUID] = uuid

	return UnmarshalledMessage{
		Database:   doc[Database].(string),
		Collection: doc[Collection].(string),
		UUID:       doc[UUID].(string),
		Payload:    payload,
	}, nil
	// fmt.Printf("fixed payloadUUID: %s\n", payload[UUID].(string))

}

func getDatabaseCollectionFromUnmarshalledMessage(docParam *map[string]interface{}) (string, string) {
	doc := *docParam
	return doc[Database].(string), doc[Collection].(string)
}

func validateUnmarshalledMessage(docParam *map[string]interface{}) error {
	doc := *docParam
	if doc[Database] == nil {
		return errors.New(fmt.Sprintf("%s field is mandatory: %v", Database, doc))
	}
	if doc[Collection] == nil {
		return errors.New(fmt.Sprintf("%s field is mandatory: %v", Collection, doc))
	}
	if doc[UUID] == nil {
		return errors.New(fmt.Sprintf("%s field is mandatory: %v", UUID, doc))
	}
	if doc[Payload] == nil {
		return errors.New(fmt.Sprintf("%s field is mandatory: %v", Payload, doc))
	}
	return nil
}

var allowedDatabases = []string{"log", "app_log"}

func validateAllowedDatabase(db string) error {

	if slices.Contains(allowedDatabases, db) == true {
		return nil
	} else {
		return errors.New(fmt.Sprintf("this database is not allowed: %s", db))
	}
}
