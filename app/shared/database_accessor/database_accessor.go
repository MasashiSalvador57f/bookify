package da

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

func InitDB(c *gin.Context) {
	dbUser := "bookify_dev"
	dbPass := "hogehoge"
	dbProtocol := "tcp"
	dbName := "bookify_development"
	dsn := fmt.Sprintf("%s%s@%s/%s", dbUser, dbPass, dbProtocol, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("error on connecting db", err)
	}

	dbmap := gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	c.Set("dbmap", dbmap)
}

/*
type DatabaseAccessor struct {
	Dbm gorp.DbMap
	ModelName string
}
*/
