package postgres_service

import (
	"fmt"
)

type PostgresService struct {
	DbConfigs map[string]PostgresConfig
}

type PostgresLogData struct {
	HostNickname string      `json:hostnickname`
	Database     string      `json:database`
	Table        string      `json:table`
	Timestamp    string      `json:timestamp`
	Data         interface{} `json:data`
}

func New(configFilePath *string) *PostgresService {
	ref := new(PostgresService)
	ref.Init(configFilePath)
	return ref
}

func (s *PostgresService) Init(configFilePath *string) {
	s.DbConfigs = LoadConfig(configFilePath)
	fmt.Println("Initializing PostgresService..")

	for k := range s.DbConfigs {
		fmt.Println(s.DbConfigs[k])
	}

}

func (s *PostgresService) PutData(data PostgresLogData) error {
	fmt.Println("Postgres PutData..")

	for k := range s.DbConfigs {
		fmt.Println(s.DbConfigs[k])
	}
	return nil
}
