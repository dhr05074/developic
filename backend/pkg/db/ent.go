package db

import (
	"ariga.io/entcache"
	"code-connect/pkg/log"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	l = log.NewZap().With("pkg", "db")
)

const dbMaxOpenConnections = 4

func MustCachedEntDriver() *entcache.Driver {
	db, err := sql.Open(dialect.MySQL, "root:teamgrit8266@tcp(mariadb:3306)/aiign?parseTime=true")
	if err != nil {
		l.Errorw("DB 접속에 실패했습니다.", "error", err)
		panic(err)
	}
	db.DB().SetMaxOpenConns(dbMaxOpenConnections)

	// Decorates the sql.Driver with entcache.Driver.
	drv := entcache.NewDriver(db)

	return drv
}
