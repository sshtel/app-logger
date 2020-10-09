package mongo_pool_service

type MongoLogData struct {
	HostNickname string `json:hostnickname`
	Database string 	`json:database`
	Collection string 	`json:collection`
	Timestamp string 	`json:timestamp`
	Data interface{} 	`json:data`
}
