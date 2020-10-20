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

func New() *PostgresService {
	ref := new(PostgresService)
	return ref
}

func (s *PostgresService) Init(conf map[string]PostgresConfig) {
	fmt.Println("Initializing PostgresService..")

	s.DbConfigs = conf
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
