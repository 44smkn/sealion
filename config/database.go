package config

import (
	"database/sql"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type DbConfig struct {
	User     string `default:"root"`
	Password string `default:"mochoten"`
	Host     string `default:"mysql"`
	Port     string `default:"3306"`
}

func GetDbConn() (*sql.DB, error) {
	var d DbConfig
	err := envconfig.Process("mysql", &d)
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/sealion?parseTime=true", d.User, d.Password, d.Host, d.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
