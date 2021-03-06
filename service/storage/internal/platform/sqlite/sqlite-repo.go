package sqlite

import (
	"database/sql"
	"errors"
	"log"

	"github.com/bottos-project/magiccube/service/storage/util"
	_ "github.com/mattn/go-sqlite3"

	"fmt"
)

const DefaultDbpath string = "./bottos.db"
const SqlTableExists string = "select count(*) from sqlite_master where type= 'table' and name="

type SqliteRepository struct {
}

// NewSqliteRepository creates a new SqliteRepository
func NewSqliteRepository() *SqliteRepository {
	return &SqliteRepository{}
}

type SqliteContext struct {
	db *sql.DB
}

func ConnectDB() (*SqliteContext, error) {
	db, err := sql.Open("sqlite3", DefaultDbpath)
	if err != nil {
		return nil, errors.New("ConnectDB sql open sqlite3 database failed")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.New("ConnectDB db ping failed")
	}
	return &SqliteContext{db}, nil
}
func (c *SqliteContext) IsTableExist(table string) bool {
	sql := SqlTableExists + "'" + table + "';"
	var num uint32
	rows, err := c.db.Query(sql)
	if err != nil {
		fmt.Println(sql)
		log.Println("%s db query failed", util.FuncLog())
		return false
	}
	//select count(*) from sqlite_master where type= 'table' and name='sync';
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			log.Println("%s", util.FuncLog(), err)
			return false
		}
		if num == 0 {
			return false
		}
	}
	return true
}
