package model

type LogData struct {
	Timestamp string `json:timestamp`
	Body interface{} `json:body`
	QueryParam interface{} `queryParam`
}
