package postgres_service

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgreSqlConn struct {
	Host     string
	Port     string
	User     string
	Pass     string
	DbName   string
	DbClient *sql.DB

	DbClientPg *gorm.DB
}

type User struct {
	Name string
	Age  int
}

func (c *PostgreSqlConn) Open() error {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Pass, c.DbName)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlInfo,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	c.DbClientPg = db
	if err != nil {
		return err
	}
	return nil
}

func (c *PostgreSqlConn) Insert(tableName string, record interface{}) error {
	result := c.DbClientPg.Table(tableName).Create(record)
	fmt.Println(result)
	return nil
}

func (c *PostgreSqlConn) GetDatabase(tableName string) *gorm.DB {
	return c.DbClientPg.Table(tableName)
}

func (c *PostgreSqlConn) Update(tableName string, record interface{}) error {
	c.DbClientPg.Table(tableName).Updates(record)
	return nil
}

func (c *PostgreSqlConn) Upsert(tableName string, record interface{}) error {
	// Do nothing on conflict
	c.DbClientPg.Table(tableName).Clauses(clause.OnConflict{DoNothing: true}).Create(record)
	c.DbClientPg.Table(tableName).Updates(record)
	return nil
}

func (c *PostgreSqlConn) Close() {
	if c.DbClient != nil {
		c.DbClient.Close()
	}
}
