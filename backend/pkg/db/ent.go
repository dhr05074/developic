package db

import (
	"ariga.io/entcache"
	"code-connect/pkg/log"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

var l = log.NewZapSugaredLogger().With("pkg", "db")

func NewCachedEntDriver() *entcache.Driver {
	db, err := sql.Open(dialect.MySQL, "root:teamgrit8266@tcp(mariadb:3306)/aiign?parseTime=true")
	if err != nil {
		l.Fatalw("failed opening connection to sqlite", "err", err)
	}
	db.DB().SetMaxOpenConns(4)

	// Decorates the sql.Driver with entcache.Driver.
	drv := entcache.NewDriver(db)

	return drv
}
