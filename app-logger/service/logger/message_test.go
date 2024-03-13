package service_logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unmarshalBaseMessage(t *testing.T) {
	jsonString := `{"database":"not_allowed_db","collection":"coll","uuid":"uuid1","payload":{"ordertime":1497014222380,"orderid":18,"itemid":"Item_184","address":{"city":"Mountain View","state":"CA","zipcode":94041}}}`
	_, err := unmarshalBaseMessage([]byte(jsonString))
	assert.Error(t, err)

	jsonString = `{"database":"log","collection":"coll","uuid":"uuid1","payload":{"ordertime":1497014222380,"orderid":18,"itemid":"Item_184","address":{"city":"Mountain View","state":"CA","zipcode":94041}}}`
	unmarshalledDoc, err := unmarshalBaseMessage([]byte(jsonString))
	assert.NoError(t, err)
	assert.NotNil(t, unmarshalledDoc)
	assert.Equal(t, unmarshalledDoc.Database, "log")
	// payload's uuid is filled as UUID
	assert.Equal(t, unmarshalledDoc.Payload["uuid"], unmarshalledDoc.UUID)

	// database field is missing
	jsonString = `{"collection":"coll","uuid":"uuid1","payload":{"ordertime":1497014222380,"orderid":18,"itemid":"Item_184","address":{"city":"Mountain View","state":"CA","zipcode":94041}}}`
	_, err = unmarshalBaseMessage([]byte(jsonString))
	assert.Error(t, err)

	// collection field is missing
	jsonString = `{"database":"log","uuid":"uuid1","payload":{"ordertime":1497014222380,"orderid":18,"itemid":"Item_184","address":{"city":"Mountain View","state":"CA","zipcode":94041}}}`
	_, err = unmarshalBaseMessage([]byte(jsonString))
	assert.Error(t, err)

	//uuid field is missing
	jsonString = `{"database":"log","collection":"coll","payload":{"ordertime":1497014222380,"orderid":18,"itemid":"Item_184","address":{"city":"Mountain View","state":"CA","zipcode":94041}}}`
	_, err = unmarshalBaseMessage([]byte(jsonString))
	assert.Error(t, err)

	//payload field is missing
	jsonString = `{"database":"log","collection":"coll","uuid":"uuid111"}`
	_, err = unmarshalBaseMessage([]byte(jsonString))
	assert.Error(t, err)

}

func Test_validateAllowedDatabase(t *testing.T) {
	assert.NoError(t, validateAllowedDatabase("log"))
	assert.NoError(t, validateAllowedDatabase("app_log"))
	assert.Error(t, validateAllowedDatabase("not_allowed_db"))

}
