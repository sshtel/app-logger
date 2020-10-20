package mysql_service

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlConn struct {
	Host     string
	Port     string
	User     string
	Pass     string
	DbName   string
	DbClient *sql.DB
}

func (c *MySqlConn) Open() error {

	target := c.User + ":" + c.Pass + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.DbName
	log.Println("Open DB connection to " + target)

	db, err := sql.Open("mysql", target)
	if err != nil {
		log.Println(err)
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	c.DbClient = db
	return nil
}

func (c *MySqlConn) MakeInsertQuery(tableName string, columns []string, values []string) string {
	query := `INSERT INTO ` + tableName
	query += "(" + strings.Join(columns, ",") + ")\n"
	query += "VALUES\n"
	query += "(" + strings.Join(values, ",") + ")"
	return query
}

func (c *MySqlConn) MakeInsertQueryFromMap(tableName string, dataMap map[string]string) string {
	columns := []string{}
	values := []string{}
	for k, v := range dataMap {
		columns = append(columns, k)
		values = append(values, v)
	}

	return c.MakeInsertQuery(tableName, columns, values)
}

func (c *MySqlConn) MakeUpdateQuery(tableName string, dataMap map[string]string, conditionKey string, conditionValue string) string {
	query := `UPDATE ` + tableName + "\n"
	query += "SET "

	sets := []string{}

	for k, v := range dataMap {
		sets = append(sets, k+"="+v)
	}

	query += strings.Join(sets, ",")

	query += "\nWHERE " + conditionKey + "=" + conditionValue
	return query
}

func (c *MySqlConn) Insert(tableName string, dataMap map[string]string) error {
	query := c.MakeInsertQueryFromMap(tableName, dataMap)

	TEST_RECORD_ONLY := os.Getenv("TEST_RECORD_ONLY")
	if TEST_RECORD_ONLY == "true" {
		log.Println(query)
		return nil
	}

	if c.DbClient == nil {
		return nil
	}

	insert, err := c.DbClient.Query(query)
	if err != nil {
		return err
	}
	insert.Close()
	return nil
}

func (c *MySqlConn) Update(tableName string, dataMap map[string]string, conditionKey string, conditionValue string) error {
	query := c.MakeUpdateQuery(tableName, dataMap, conditionKey, conditionValue)

	TEST_RECORD_ONLY := os.Getenv("TEST_RECORD_ONLY")
	if TEST_RECORD_ONLY == "true" {
		log.Println(query)
		return nil
	}

	if c.DbClient == nil {
		return nil
	}

	insert, err := c.DbClient.Query(query)
	if err != nil {
		return err
	}
	insert.Close()
	return nil
}

func (c *MySqlConn) Close() {
	if c.DbClient != nil {
		c.DbClient.Close()
	}
}
