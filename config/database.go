package config

import (
	"database/sql"
	"fmt"

	"github.com/google/wire"

	"github.com/kelseyhightower/envconfig"
)

var Set = wire.NewSet(GetDbConn)

type DbConfig struct {
	User     string `default:"root"`
	Password string `default:"mochoten"`
	Host     string `default:"mysql"`
	Port     string `default:"3306"`
}

func GetDbConn() *sql.DB {
	var d DbConfig
	err := envconfig.Process("mysql", &d)
	if err != nil {
		// TODO: error handling
		return nil
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/sealion?parseTime=true", d.User, d.Password, d.Host, d.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// TODO: error handling
		return nil
	}
	return db
}
